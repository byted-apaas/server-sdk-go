package utils

import (
	"context"
	"reflect"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

// GetParamUnauthField 获取参数的所有无权限字段
func GetParamUnauthField(ctx context.Context) map[string]interface{} {
	return cUtils.GetParamUnauthFieldMapFromCtx(ctx)
}

// GetParamUnauthFieldByKey 通过 key 获取参数的无权限字段
// - 如果参数类型为 Record 则返回 []string
// - 如果参数类型为 RecordList，则返回 [][]string
func GetParamUnauthFieldByKey(ctx context.Context, key string) interface{} {
	unauthFields, _ := cUtils.GetParamUnauthFieldMapFromCtx(ctx)[key]
	return unauthFields
}

// GetParamUnauthFieldRecordByKey 通过 key 获取记录类型参数的无权限字段
func GetParamUnauthFieldRecordByKey(ctx context.Context, key string) (unauthFields []string) {
	return cUtils.GetParamUnauthFieldRecordByKey(ctx, key)
}

// GetParamUnauthFieldRecordListByKey 通过 key 获取记录列表类型参数的无权限字段
func GetParamUnauthFieldRecordListByKey(ctx context.Context, key string) (unauthFieldsList [][]string) {
	return cUtils.GetParamUnauthFieldRecordListByKey(ctx, key)
}

func GetRecordUnauthFieldByObject(ctx context.Context, objectAPIName string) map[int64][]string {
	return cUtils.GetRecordUnauthFieldByObject(ctx, objectAPIName)
}

func GetRecordUnauthFieldByObjectAndRecordID(ctx context.Context, objectAPIName string, recordID int64) []string {
	return cUtils.GetRecordUnauthFieldByObjectAndRecordID(ctx, objectAPIName, recordID)
}

func ParseUnauthFields(unauthFieldInput [][]string, unauthFieldOutput interface{}) error {
	// 不接收权限信息或无权限信息，直接返回
	if len(unauthFieldInput) == 0 {
		return nil
	}

	// 接收参数类型校验
	_, ok1 := unauthFieldOutput.(*[][]string) // 批量查询场景
	_, ok2 := unauthFieldOutput.(*[]string)   // 单个查询场景
	if !ok1 && !ok2 {
		return cExceptions.InvalidParamError("the type of unauthFields should be *[][]string or *[]string, but %T", unauthFieldOutput)
	}

	// 权限信息写入接收参数
	if ok1 {
		// 批量查询场景
		refValues := reflect.ValueOf(unauthFieldOutput).Elem()
		for i := range unauthFieldInput {
			refValues.Set(reflect.Append(refValues, reflect.ValueOf(unauthFieldInput[i])))
		}
	} else {
		// 单个查询场景
		refValues := reflect.ValueOf(unauthFieldOutput).Elem()
		for _, v := range unauthFieldInput[0] {
			refValues.Set(reflect.Append(refValues, reflect.ValueOf(v)))
		}
	}

	return nil
}