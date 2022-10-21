// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/data"
	cond2 "github.com/byted-apaas/server-sdk-go/service/data/cond"
	"github.com/byted-apaas/server-sdk-go/service/data/op"
)

type Query struct {
	appCtx        *structs.AppCtx
	objectAPIName string
	limit         int64
	offset        int64
	fields        []string
	order         []*structs.Order
	filter        *cond2.LogicalExpression
	err           error
}

func (q *Query) Count(ctx context.Context) (int64, error) {
	if q.appCtx.IsOpenSDK() {
		param := &structs.GetRecordsReqParamV2{
			Limit:  1,
			Offset: q.offset,
			Fields: q.fields,
			Sort:   q.order,
			Count:  cUtils.BoolPtr(true),
			Filter: q.buildCriterionV2(q.filter),
		}

		return request.GetInstance(ctx).GetRecordCountV2(ctx, q.appCtx, q.objectAPIName, param)
	} else {
		param := &structs.GetRecordsReqParam{
			Limit:          1,
			Offset:         q.offset,
			FieldApiNames:  q.fields,
			Order:          q.order,
			NeedTotalCount: true,
		}

		if criterion, err := q.buildCriterion(ctx, q.filter); err != nil {
			return 0, err
		} else {
			param.Criterion = criterion
		}

		return request.GetInstance(ctx).GetRecordCount(ctx, q.appCtx, q.objectAPIName, param)
	}
}

func (q *Query) FindAll(ctx context.Context, records interface{}) error {
	if q.appCtx.IsOpenSDK() {
		return q.findAllV2(ctx, records)
	}
	return q.findAll(ctx, records)
}

func (q *Query) Find(ctx context.Context, records interface{}) error {
	if q.err != nil {
		return q.err
	}

	// OpenSDK 走新接口，因为新接口有权限控制
	if q.appCtx.IsOpenSDK() {
		param := &structs.GetRecordsReqParamV2{
			Limit:  q.limit,
			Offset: q.offset,
			Fields: q.fields,
			Filter: q.buildCriterionV2(q.filter),
			Sort:   q.order,
			Count:  cUtils.BoolPtr(false),
		}

		return request.GetInstance(ctx).GetRecordsV2(ctx, q.appCtx, q.objectAPIName, param, records)
	} else {
		criterion, err := q.buildCriterion(ctx, q.filter)
		if err != nil {
			return err
		}

		param := &structs.GetRecordsReqParam{
			Limit:          q.limit,
			Offset:         q.offset,
			FieldApiNames:  q.fields,
			Order:          q.order,
			NeedTotalCount: false,
			Criterion:      criterion,
		}

		return request.GetInstance(ctx).GetRecords(ctx, q.appCtx, q.objectAPIName, param, records)
	}
}

func (q *Query) FindOne(ctx context.Context, record interface{}) error {
	if q.err != nil {
		return q.err
	}

	var (
		records []interface{}
		err     error
	)

	// OpenSDK 走新接口，因为新接口有权限控制
	if q.appCtx.IsOpenSDK() {
		param := &structs.GetRecordsReqParamV2{
			Limit:  1,
			Offset: q.offset,
			Fields: q.fields,
			Filter: q.buildCriterionV2(q.filter),
			Sort:   q.order,
			Count:  cUtils.BoolPtr(false),
		}

		err = request.GetInstance(ctx).GetRecordsV2(ctx, q.appCtx, q.objectAPIName, param, &records)
	} else {
		param := &structs.GetRecordsReqParam{
			Limit:          1,
			Offset:         q.offset,
			FieldApiNames:  q.fields,
			Order:          q.order,
			NeedTotalCount: false,
		}

		if criterion, e := q.buildCriterion(ctx, q.filter); e != nil {
			return e
		} else {
			param.Criterion = criterion
		}

		err = request.GetInstance(ctx).GetRecords(ctx, q.appCtx, q.objectAPIName, param, &records)
	}

	if err != nil {
		return err
	}

	if len(records) < 1 {
		return cExceptions.InvalidParamError("record not found")
	}

	err = cUtils.Decode(records[0], record)
	if err != nil {
		return cExceptions.InternalError("Decode failed: %+v", err)
	}
	return nil
}

func newQuery(s *structs.AppCtx, objectAPIName string, err error) *Query {
	q := &Query{
		appCtx:        s,
		objectAPIName: objectAPIName,
		limit:         constants.PageLimitMax,
		offset:        0,
		fields:        []string{},
		order:         []*structs.Order{},
		filter:        cond2.NewLogicalExpression(op.And, nil, nil),
	}

	if err != nil {
		q.err = err
		return q
	}

	if objectAPIName == "" {
		q.err = cExceptions.InvalidParamError("objectAPIName is empty")
		return q
	}

	return q
}

// Where 配置过滤条件
// @param condition：过滤条件，其类型为 *cond.LogicalExpression 或 *cond.ArithmeticExpression，不合法的类型会报错
// @example condition：
//     cond.And(...)
//     cond.Or(...)
//     cond.Eq(...)
//     cond.Gt(...)
// @return 返回查询对象
func (q *Query) Where(condition interface{}) data.IQuery {
	if q.err != nil {
		return q
	}

	if condition == nil {
		return q
	}

	switch condition.(type) {
	case cond2.LogicalExpression:
		v, _ := condition.(cond2.LogicalExpression)
		q.filter.AddLogicalExpression(&v)
	case *cond2.LogicalExpression:
		v, _ := condition.(*cond2.LogicalExpression)
		q.filter.AddLogicalExpression(v)
	case cond2.ArithmeticExpression:
		v, _ := condition.(cond2.ArithmeticExpression)
		q.filter.AddArithmeticExpression(&v)
	case *cond2.ArithmeticExpression:
		v, _ := condition.(*cond2.ArithmeticExpression)
		q.filter.AddArithmeticExpression(v)
	default:
		q.err = cExceptions.InvalidParamError("Query.Where received invalid type, should be *cond.LogicalExpression or *cond.ArithmeticExpression, but received %s ", reflect.TypeOf(condition))
	}
	return q
}

func (q *Query) Offset(offset int64) data.IQuery {
	if q.err != nil {
		return q
	}

	if offset < 0 {
		q.err = cExceptions.InvalidParamError("Query.Offset received invalid value, should >= 0")
	}

	q.offset = offset
	return q
}

func (q *Query) Limit(limit int64) data.IQuery {
	if q.err != nil {
		return q
	}

	if limit < 1 || limit > constants.PageLimitMax {
		q.err = cExceptions.InvalidParamError("Query.Limit received invalid value (%d), should be 1~%d", limit, constants.PageLimitMax)
	}

	q.limit = limit
	return q
}

func (q *Query) OrderBy(fieldAPINames ...string) data.IQuery {
	if q.err != nil {
		return q
	}

	for _, fieldAPIName := range fieldAPINames {
		q.order = append(q.order, &structs.Order{
			Field:     fieldAPIName,
			Direction: constants.OrderAsc,
		})
	}
	return q
}

func (q *Query) OrderByDesc(fieldAPINames ...string) data.IQuery {
	if q.err != nil {
		return q
	}

	for _, fieldAPIName := range fieldAPINames {
		q.order = append(q.order, &structs.Order{
			Field:     fieldAPIName,
			Direction: constants.OrderDesc,
		})
	}
	return q
}

func (q *Query) Select(fieldAPINames ...string) data.IQuery {
	if q.err != nil {
		return q
	}

	q.fields = append(q.fields, fieldAPINames...)
	return q
}

// buildCriterion 构建查询过滤条件
func (q *Query) buildCriterion(ctx context.Context, filter *cond2.LogicalExpression) (*cond2.Criterion, error) {
	if q.err != nil {
		return nil, q.err
	}

	if filter == nil {
		return nil, nil
	}

	// 广度优先遍历将树状过滤条件拍平
	conditions, indexToFieldPaths, hasIsDeleted := q.bfs(filter)
	// 深度优先遍历根据树状过滤条件生成逻辑表达式
	logic := q.dfs(filter)

	err := q.fillFieldPaths(ctx, indexToFieldPaths)
	if err != nil {
		return nil, err
	}

	// 补充 _isDeleted 条件
	if !hasIsDeleted {
		if len(logic) > 0 {
			logic = fmt.Sprintf("(%d and %s)", len(conditions)+1, logic)
		} else {
			logic = fmt.Sprintf("%d", len(conditions)+1)
		}
		isDeletedCondition := cond2.Eq("_isDeleted", false)
		isDeletedCondition.Index = int64(len(conditions) + 1)
		isDeletedCondition.Left.Settings.FieldPath[0].ObjectAPIName = q.objectAPIName
		conditions = append(conditions, isDeletedCondition)
	}
	return &cond2.Criterion{Conditions: conditions, Logic: logic}, nil
}

// bfs 广度优先遍历
//     1.设置表达式的序号
//     2.设置第一层的 ObjectAPIName
//     3.记录多层下钻场景的 ObjectAPIName，后绪统一补全，减少 getFields 的次数
//     4.判断是否有 _isDeleted 条件，如果没有，后绪需要补上
// @param filter：查询的过滤条件
// @return1：拍平后的条件
// @return2：多层下钻场景有 ObjectAPIName 待补全
// @return3：是否有 _isDeleted 条件
func (q *Query) bfs(filter *cond2.LogicalExpression) ([]*cond2.ArithmeticExpression, map[int64][]*cond2.FieldPath, bool) {
	if filter == nil {
		return nil, nil, false
	}

	var (
		queue             = []interface{}{filter}
		sequence          = int64(1)
		conditions        []*cond2.ArithmeticExpression
		indexToFieldPaths = map[int64][]*cond2.FieldPath{}
		hasIsDeleted      = false
	)

	for len(queue) > 0 {
		// 出队
		ele := queue[0]
		queue = queue[1:]

		// 如果是算术表达式
		if v, ok := ele.(*cond2.ArithmeticExpression); ok {
			// 1.设置表达式的序号
			v.Index, sequence = sequence, sequence+1
			conditions = append(conditions, v)

			// 2.设置第一层的 ObjectAPIName
			v.Left.Settings.FieldPath[0].ObjectAPIName = q.objectAPIName

			// 3.记录多层下钻场景的 ObjectAPIName，后绪统一补全，减少 getFields 的次数
			if len(v.Left.Settings.FieldPath) > 1 {
				indexToFieldPaths[v.Index] = v.Left.Settings.FieldPath
			}

			// 4.判断是否有 _isDeleted 条件，如果没有，后绪需要补上
			if v.Left.Settings.FieldPath[0].FieldAPIName == "_isDeleted" {
				hasIsDeleted = true
			}
			continue
		}

		// 如果是逻辑表达式
		if v, ok := ele.(*cond2.LogicalExpression); ok {
			for i := range v.LogicalExpressions {
				// 进队
				queue = append(queue, v.LogicalExpressions[i])
			}
			for i := range v.ArithmeticExpressions {
				// 进队
				queue = append(queue, v.ArithmeticExpressions[i])
			}
			continue
		}
	}

	return conditions, indexToFieldPaths, hasIsDeleted
}

// dfs 深度优先遍历，依据树状过滤条件生成逻辑表达式
func (q *Query) dfs(filter *cond2.LogicalExpression) string {
	if filter == nil || (len(filter.LogicalExpressions) == 0 && len(filter.ArithmeticExpressions) == 0) {
		return ""
	}

	var logics []string
	for i := range filter.LogicalExpressions {
		logics = append(logics, q.dfs(filter.LogicalExpressions[i]))
	}

	for i := range filter.ArithmeticExpressions {
		logics = append(logics, strconv.FormatInt(filter.ArithmeticExpressions[i].Index, 10))
	}
	return "(" + strings.Join(logics, fmt.Sprintf(" %s ", filter.Type)) + ")"
}

func (q *Query) fillFieldPaths(ctx context.Context, indexToFieldPaths map[int64][]*cond2.FieldPath) error {
	// 计算最大下钻层数
	maxLayer := 0
	for _, fieldPaths := range indexToFieldPaths {
		if maxLayer < len(fieldPaths) {
			maxLayer = len(fieldPaths)
		}
	}

	// objectAPIName.fieldAPIName: lookupObjectAPIName
	objectFieldToLookupObjectAPIName := map[string]string{}
	for i := 0; i < maxLayer; i++ {
		// objectAPIName: fieldAPINames，实现去重合并，避免重复查询
		objectToFieldAPINames := map[string]map[string]string{}
		for _, fieldPaths := range indexToFieldPaths {
			if len(fieldPaths) < i+1 {
				continue
			}

			if fieldPaths[i].ObjectAPIName != "" {
				objectAPIName := fieldPaths[i].ObjectAPIName
				fieldAPIName := fieldPaths[i].FieldAPIName
				key := fmt.Sprintf("%s.%s", objectAPIName, fieldAPIName)
				if lookupObjectAPIName, ok := objectFieldToLookupObjectAPIName[key]; !ok || lookupObjectAPIName == "" {
					// 提取 FieldPath
					fieldAPINames, _ := objectToFieldAPINames[objectAPIName]
					if fieldAPINames == nil {
						fieldAPINames = map[string]string{}
					}
					fieldAPINames[fieldAPIName] = ""
					objectToFieldAPINames[objectAPIName] = fieldAPINames
				}
			} else {
				// 需要补全时，上一层一定有值
				if i-1 < 0 {
					return cExceptions.InternalError("fillFieldPaths failed, reason: i-1<0, fieldPath: %+v", fieldPaths[i])
				}

				key := fmt.Sprintf("%s.%s", fieldPaths[i-1].ObjectAPIName, fieldPaths[i-1].FieldAPIName)
				if lookupObjectAPIName, ok := objectFieldToLookupObjectAPIName[key]; !ok || lookupObjectAPIName == "" {
					return cExceptions.InternalError("fillFieldPaths failed, key: %s, lookupObjectAPIName: %+v", key, objectFieldToLookupObjectAPIName)
				} else {
					// 补全 lookupObjectAPIName
					fieldPaths[i].ObjectAPIName = lookupObjectAPIName
					// 添加到 map 中，为下一步补全做准备
					key := fmt.Sprintf("%s.%s", fieldPaths[i].ObjectAPIName, fieldPaths[i].FieldAPIName)
					if lookupObjectAPIName, ok := objectFieldToLookupObjectAPIName[key]; (!ok || lookupObjectAPIName == "") && i < maxLayer {
						objectAPIName := fieldPaths[i].ObjectAPIName
						fieldAPIName := fieldPaths[i].FieldAPIName
						fieldAPINames, _ := objectToFieldAPINames[objectAPIName]
						if fieldAPINames == nil {
							fieldAPINames = map[string]string{}
						}
						fieldAPINames[fieldAPIName] = ""
						objectToFieldAPINames[objectAPIName] = fieldAPINames
					}
				}
			}
		}

		for objectAPIName, fieldAPINames := range objectToFieldAPINames {
			var fields []string
			for fieldAPIName := range fieldAPINames {
				fields = append(fields, fieldAPIName)
			}
			apiNameMap, err := q.getLookupObjectAPINames(ctx, objectAPIName, fields...)
			if err != nil {
				return err
			}
			for key, value := range apiNameMap {
				objectFieldToLookupObjectAPIName[key] = value
			}
		}
	}
	return nil
}

type FieldInfo struct {
	APIName string `json:"api_name"`
	Type    struct {
		Settings struct {
			LookupObjectAPIName string `json:"referenced_object_api_name"`
		} `json:"settings"`
	} `json:"type"`
}

type ObjectInfo struct {
	Fields []*FieldInfo `json:"fields"`
}

func (q *Query) getLookupObjectAPINames(ctx context.Context, objectAPIName string, fieldAPINames ...string) (map[string]string, error) {
	if len(fieldAPINames) == 0 {
		return nil, nil
	}

	if len(fieldAPINames) == 1 {
		field, err := request.GetInstance(ctx).GetField(ctx, q.appCtx, objectAPIName, fieldAPINames[0])
		if err != nil {
			return nil, err
		}

		if field == nil {
			return nil, cExceptions.InvalidParamError("Relate field (%s) is not exist", fieldAPINames[0])
		}

		var fieldInfo FieldInfo
		err = cUtils.Decode(field, &fieldInfo)
		if err != nil {
			return nil, cExceptions.InternalError("Decode fieldInfo failed, err: %+v", err)
		}

		return map[string]string{
			fmt.Sprintf("%s.%s", objectAPIName, fieldAPINames[0]): fieldInfo.Type.Settings.LookupObjectAPIName,
		}, nil
	}

	fields, err := request.GetInstance(ctx).GetFields(ctx, q.appCtx, objectAPIName)
	if err != nil {
		return nil, err
	}

	var objectInfo ObjectInfo
	err = cUtils.Decode(fields, &objectInfo)
	if err != nil {
		return nil, cExceptions.InternalError("Decode ObjectInfo failed, err: %+v", err)
	}

	objectFieldToLookObjectAPIName := map[string]string{}
	for _, field := range objectInfo.Fields {
		if cUtils.StrInStrs(fieldAPINames, field.APIName) {
			objectFieldToLookObjectAPIName[fmt.Sprintf("%s.%s", objectAPIName, field.APIName)] = field.Type.Settings.LookupObjectAPIName
		}
	}
	return objectFieldToLookObjectAPIName, nil
}

func (q *Query) buildCriterionV2(filter *cond2.LogicalExpression) []interface{} {
	if filter == nil || (len(filter.LogicalExpressions) == 0 && len(filter.ArithmeticExpressions) == 0) {
		return nil
	}

	var conditions []interface{}
	for i := range filter.LogicalExpressions {
		conditions = append(conditions, map[string]interface{}{
			filter.LogicalExpressions[i].Type: q.buildCriterionV2(filter.LogicalExpressions[i]),
		})
	}

	for i := range filter.ArithmeticExpressions {
		conditions = append(conditions, filter.ArithmeticExpressions[i].Expr)
	}
	return conditions
}

func (q *Query) findAll(ctx context.Context, records interface{}) error {
	var (
		totalRecords []map[string]interface{}
		maxID        = int64(0)
	)

	criterion, err := q.buildCriterion(ctx, q.filter)
	if err != nil {
		return err
	}

	// 添加 _id > maxID 条件
	if len(criterion.Logic) > 0 {
		criterion.Logic = fmt.Sprintf("(%d and %s)", len(criterion.Conditions)+1, criterion.Logic)
	} else {
		criterion.Logic = fmt.Sprintf("%d", len(criterion.Conditions)+1)
	}
	idCondition := cond2.Gt("_id", maxID)
	idCondition.Index = int64(len(criterion.Conditions) + 1)
	idCondition.Left.Settings.FieldPath[0].ObjectAPIName = q.objectAPIName
	criterion.Conditions = append(criterion.Conditions, idCondition)

	param := &structs.GetRecordsReqParam{
		Limit:          constants.PageLimitMax,
		Offset:         0,
		FieldApiNames:  q.fields,
		Criterion:      criterion,
		Order:          []*structs.Order{{Field: "_id", Direction: "asc"}},
		NeedTotalCount: false,
	}

	for {
		var perRecords []map[string]interface{}
		idCondition.Right.Settings.Data = maxID
		err := request.GetInstance(ctx).GetRecords(ctx, q.appCtx, q.objectAPIName, param, &perRecords)
		if err != nil {
			return err
		}
		totalRecords = append(totalRecords, perRecords...)

		// 更新 maxID
		for _, perRecord := range perRecords {
			id, err := utils.ParseInt64(perRecord["_id"])
			if err != nil {
				return err
			}
			if id > maxID {
				maxID = id
			}
		}

		if len(perRecords) < constants.PageLimitMax {
			break
		}
	}

	marshal, err := cUtils.JsonMarshalBytes(totalRecords)
	if err != nil {
		return cExceptions.InternalError("json.Marshal failed: %+v", err)
	}

	err = cUtils.JsonUnmarshalBytes(marshal, records)
	if err != nil {
		return cExceptions.InvalidParamError("json.Marshal failed: %+v", err)
	}
	return nil
}

func (q *Query) findAllV2(ctx context.Context, records interface{}) error {
	var (
		totalRecords []map[string]interface{}
		maxID        = int64(0)
	)

	idCondition := cond2.NewExpressionV2("_id", op.Gt, maxID)
	conditions := q.buildCriterionV2(q.filter)
	conditions = append(conditions, idCondition)

	param := &structs.GetRecordsReqParamV2{
		Limit:  constants.PageLimitMax,
		Offset: 0,
		Fields: q.fields,
		Sort:   []*structs.Order{{Field: "_id", Direction: "asc"}},
		Count:  cUtils.BoolPtr(false),
		Filter: conditions,
	}

	for {
		var perRecords []map[string]interface{}
		idCondition.RightValue = maxID
		err := request.GetInstance(ctx).GetRecordsV2(ctx, q.appCtx, q.objectAPIName, param, &perRecords)
		if err != nil {
			return err
		}
		totalRecords = append(totalRecords, perRecords...)

		// 更新 maxID
		for _, perRecord := range perRecords {
			id, err := utils.ParseInt64(perRecord["_id"])
			if err != nil {
				return err
			}
			if id > maxID {
				maxID = id
			}
		}

		if len(perRecords) < constants.PageLimitMax {
			break
		}
	}

	marshal, err := cUtils.JsonMarshalBytes(totalRecords)
	if err != nil {
		return cExceptions.InternalError("json.Marshal failed: %+v", err)
	}

	err = cUtils.JsonUnmarshalBytes(marshal, records)
	if err != nil {
		return cExceptions.InvalidParamError("json.Marshal failed: %+v", err)
	}
	return nil
}
