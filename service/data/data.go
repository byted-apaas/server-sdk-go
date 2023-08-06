// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"
	"reflect"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

//go:generate mockery --name=IData --structname=Data --filename=Data.go
type IData interface {
	Object(objectAPIName string) IObject
	NewTransaction() ITransaction
	Oql(oql string, args ...interface{}) IOql
}

//go:generate mockery --name=IDataV2 --structname=DataV2 --filename=DataV2.go
type IDataV2 interface {
	Object(objectAPIName string) IObjectV2
}

//go:generate mockery --name=IObject --structname=Object --filename=Object.go
type IObject interface {
	Create(ctx context.Context, record interface{}) (id *structs.RecordID, err error)
	BatchCreate(ctx context.Context, records interface{}) (ids []int64, err error)
	BatchCreateAsync(ctx context.Context, records interface{}) (taskID int64, err error)

	Update(ctx context.Context, recordID int64, record interface{}) (err error)
	BatchUpdate(ctx context.Context, records map[int64]interface{}, result ...interface{}) (err error)
	BatchUpdateAsync(ctx context.Context, records map[int64]interface{}) (taskID int64, err error)

	Delete(ctx context.Context, recordID int64) (err error)
	BatchDelete(ctx context.Context, recordIDs []int64, result ...interface{}) (err error)
	BatchDeleteAsync(ctx context.Context, recordIDs []int64) (taskID int64, err error)

	Count(ctx context.Context) (count int64, err error)
	Find(ctx context.Context, records interface{}, unauthFields ...interface{}) (err error)
	FindOne(ctx context.Context, record interface{}, unauthFields ...interface{}) (err error)
	// Deprecated: Use FindStream instead.
	FindAll(ctx context.Context, records interface{}) (err error)
	// FindStream 流式查询
	// @example:
	//   FindStream(ctx, reflect.TypeOf(&TestObject{}), func(ctx context.Context, records interface{}) error {
	//      var rs []*TestObject
	//		for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
	//			o, ok := reflect.ValueOf(records).Elem().Index(i).Interface().(TestObject)
	//			if !ok {
	//				panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(records).Elem().Index(i).Interface()))
	//			}
	//			rs = append(rs, &o)
	//		}
	//
	//		// doSomething
	//  })
	FindStream(ctx context.Context, recordType reflect.Type, handler func(ctx context.Context, records interface{}, unauthFields interface{}) error) (err error)

	// Where 配置过滤条件
	// @param condition：过滤条件，其类型为逻辑表达式 *cond.LogicalExpression 或算术表达式 *cond.ArithmeticExpression，不合法的类型会报错
	// @example condition：
	//     cond.And(...)
	//     cond.Or(...)
	//     cond.Eq(...)
	//     cond.Gt(...)
	// @return 返回查询对象
	Where(condition interface{}) IQuery
	// FuzzySearch 模糊查询：与 where 之间是与关系
	// @param keyword 模糊查询的关键字，必填且不可以为空串
	// @param fieldAPINames 『可搜索字段』的字段列表，不可为空
	// @example: FuzzySearch("张三", []string{"_name"})
	FuzzySearch(keyword string, fieldAPINames []string) IQuery
	Offset(offset int64) IQuery
	Limit(limit int64) IQuery
	OrderBy(fieldAPINames ...string) IQuery
	OrderByDesc(fieldAPINames ...string) IQuery
	Select(fieldAPINames ...string) IQuery

	UseUserAuth() IObject
	UseSystemAuth() IObject
}

//go:generate mockery --name=IObjectV2 --structname=ObjectV2 --filename=ObjectV2.go
type IObjectV2 interface {
	Create(ctx context.Context, record interface{}) (*structs.RecordID, error)
	BatchCreate(ctx context.Context, records interface{}) ([]int64, error)

	Update(ctx context.Context, recordID int64, record interface{}) error
	BatchUpdate(ctx context.Context, records map[int64]interface{}, result ...interface{}) error

	Delete(ctx context.Context, recordID int64) error
	BatchDelete(ctx context.Context, recordIDs []int64, result ...interface{}) error

	Count(ctx context.Context) (int64, error)
	// Find 查询记录，最多 200条
	// @param records 返回的记录
	// @param unauthFields 返回的无权限字段，类型为 [][]string，第一维对应 records 中的第 x 个记录，第二维表示无权限字段的列表
	Find(ctx context.Context, records interface{}, unauthFields ...interface{}) error
	// FindOne 查询1条记录
	// @param record 返回的记录
	// @param unauthFields 返回的无权限字段，类型为 []string，表示无权限字段的列表
	FindOne(ctx context.Context, record interface{}, unauthFields ...interface{}) error
	// Deprecated: Use FindStream instead.
	FindAll(ctx context.Context, records interface{}) error
	// FindStream 流式查询
	// @example:
	//   FindStream(ctx, reflect.TypeOf(&TestObject{}), func(ctx context.Context, records interface{}) error {
	//      var rs []*TestObject
	//		for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
	//			o, ok := reflect.ValueOf(records).Elem().Index(i).Interface().(TestObject)
	//			if !ok {
	//				panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(records).Elem().Index(i).Interface()))
	//			}
	//			rs = append(rs, &o)
	//		}
	//
	//		// doSomething
	//  })
	FindStream(ctx context.Context, recordType reflect.Type, handler func(ctx context.Context, records interface{}, unauthFields interface{}) error) error

	// Where 配置过滤条件
	// @param condition：过滤条件，其类型为逻辑表达式 *cond.LogicalExpression 或算术表达式 *cond.ArithmeticExpression，不合法的类型会报错
	// @example condition：
	//     cond.And(...)
	//     cond.Or(...)
	//     cond.Eq(...)
	//     cond.Gt(...)
	// @return 返回查询对象
	Where(condition interface{}) IQuery
	Offset(offset int64) IQuery
	Limit(limit int64) IQuery
	OrderBy(fieldAPINames ...string) IQuery
	OrderByDesc(fieldAPINames ...string) IQuery
	Select(fieldAPINames ...string) IQuery
}

//go:generate mockery --name=ITransaction --structname=Transaction --filename=Transaction.go
type ITransaction interface {
	Object(objectAPIName string) ITransactionObject
	Commit(ctx context.Context) (err error)

	UseUserAuth() ITransaction
	UseSystemAuth() ITransaction
}

//go:generate mockery --name=ITransactionObject --structname=TransactionObject --filename=TransactionObject.go
type ITransactionObject interface {
	RegisterCreate(record interface{}) (id *structs.TransactionRecordID, err error)
	RegisterUpdate(recordID interface{}, record interface{})
	RegisterDelete(recordID interface{})
	RegisterBatchCreate(records interface{}) (ids []interface{}, err error)
	RegisterBatchUpdate(records interface{})
	RegisterBatchDelete(recordIDs interface{})
}

//go:generate mockery --name=IQuery --structname=Query --filename=Query.go
type IQuery interface {
	Count(ctx context.Context) (int64, error)
	// Find 查询记录，最多 200条
	// @param records 返回的记录
	// @param unauthFields 返回的无权限字段，类型为 [][]string，第一维对应 records 中的第 x 个记录，第二维表示无权限字段的列表
	Find(ctx context.Context, records interface{}, unauthFields ...interface{}) error
	// FindOne 查询1条记录
	// @param record 返回的记录
	// @param unauthFields 返回的无权限字段，类型为 []string，表示无权限字段的列表
	FindOne(ctx context.Context, record interface{}, unauthFields ...interface{}) error
	// Deprecated: Use FindStream instead.
	FindAll(ctx context.Context, records interface{}) error
	// FindStream 流式查询
	// @example:
	//   FindStream(ctx, reflect.TypeOf(&TestObject{}), func(ctx context.Context, records interface{}) error {
	//      var rs []*TestObject
	//		for i := 0; i < reflect.ValueOf(records).Elem().Len(); i++ {
	//			o, ok := reflect.ValueOf(records).Elem().Index(i).Interface().(TestObject)
	//			if !ok {
	//				panic(fmt.Sprintf("should be TestObject, but %T", reflect.ValueOf(records).Elem().Index(i).Interface()))
	//			}
	//			rs = append(rs, &o)
	//		}
	//
	//		// doSomething
	//  })
	FindStream(ctx context.Context, recordType reflect.Type, handler func(ctx context.Context, records interface{}, unauthFields interface{}) error) error

	// Where 配置过滤条件
	// @param condition：过滤条件，其类型为 *cond.LogicalExpression 或 *cond.ArithmeticExpression，不合法的类型会报错
	// @example condition：
	//     cond.And(...)
	//     cond.Or(...)
	//     cond.Eq(...)
	//     cond.Gt(...)
	// @return 返回查询对象
	Where(condition interface{}) IQuery
	// FuzzySearch 模糊查询：与 where 之间是与关系
	// @param keyword 模糊查询的关键字，必填且不可以为空串
	// @param fieldAPINames 『可搜索字段』的字段列表，不可为空
	// @example: FuzzySearch("张三", []string{"_name"})
	FuzzySearch(keyword string, fieldAPINames []string) IQuery
	Offset(offset int64) IQuery
	Limit(limit int64) IQuery
	OrderBy(fieldAPINames ...string) IQuery
	OrderByDesc(fieldAPINames ...string) IQuery
	Select(fieldAPINames ...string) IQuery
	UseUserAuth() IQuery
	UseSystemAuth() IQuery
}

//go:generate mockery --name=IOql --structname=Oql --filename=Oql.go
type IOql interface {
	UseUserAuth() IOql
	UseSystemAuth() IOql
	Execute(ctx context.Context, resultSet interface{}, unauthFields ...interface{}) (err error)
}
