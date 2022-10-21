// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

//go:generate mockery --name=IData --structname=Data --filename=Data.go
type IData interface {
	Object(objectAPIName string) IObject
	NewTransaction() ITransaction
	Oql(oql string) IOql
}

//go:generate mockery --name=IDataV2 --structname=DataV2 --filename=DataV2.go
type IDataV2 interface {
	Object(objectAPIName string) IObjectV2
}

//go:generate mockery --name=IObject --structname=Object --filename=Object.go
type IObject interface {
	Create(ctx context.Context, record interface{}) (*structs.RecordID, error)
	BatchCreate(ctx context.Context, records interface{}) ([]int64, error)
	BatchCreateAsync(ctx context.Context, records interface{}) (int64, error)

	Update(ctx context.Context, recordID int64, record interface{}) error
	BatchUpdate(ctx context.Context, records map[int64]interface{}) error
	BatchUpdateAsync(ctx context.Context, records map[int64]interface{}) (int64, error)

	Delete(ctx context.Context, recordID int64) error
	BatchDelete(ctx context.Context, recordIDs []int64) error
	BatchDeleteAsync(ctx context.Context, recordIDs []int64) (int64, error)

	Count(ctx context.Context) (int64, error)
	Find(ctx context.Context, records interface{}) error
	FindOne(ctx context.Context, record interface{}) error
	FindAll(ctx context.Context, records interface{}) error

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

//go:generate mockery --name=IObjectV2 --structname=ObjectV2 --filename=ObjectV2.go
type IObjectV2 interface {
	Create(ctx context.Context, record interface{}) (*structs.RecordID, error)
	BatchCreate(ctx context.Context, records interface{}) ([]int64, error)

	Update(ctx context.Context, recordID int64, record interface{}) error
	BatchUpdate(ctx context.Context, records map[int64]interface{}) error

	Delete(ctx context.Context, recordID int64) error
	BatchDelete(ctx context.Context, recordIDs []int64) error

	Count(ctx context.Context) (int64, error)
	Find(ctx context.Context, records interface{}) error
	FindOne(ctx context.Context, record interface{}) error
	FindAll(ctx context.Context, records interface{}) error

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
	Commit(ctx context.Context) error
}

//go:generate mockery --name=ITransactionObject --structname=TransactionObject --filename=TransactionObject.go
type ITransactionObject interface {
	RegisterCreate(record interface{}) (*structs.TransactionRecordID, error)
	RegisterUpdate(recordID interface{}, record interface{})
	RegisterDelete(recordID interface{})
	RegisterBatchCreate(records interface{}) ([]interface{}, error)
	RegisterBatchUpdate(records interface{})
	RegisterBatchDelete(recordIDs interface{})
}

//go:generate mockery --name=IQuery --structname=Query --filename=Query.go
type IQuery interface {
	Count(ctx context.Context) (int64, error)
	Find(ctx context.Context, records interface{}) error
	FindOne(ctx context.Context, record interface{}) error
	FindAll(ctx context.Context, records interface{}) error

	// Where 配置过滤条件
	// @param condition：过滤条件，其类型为 *cond.LogicalExpression 或 *cond.ArithmeticExpression，不合法的类型会报错
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

//go:generate mockery --name=IOql --structname=Oql --filename=Oql.go
type IOql interface {
	Execute(ctx context.Context, resultSet interface{}) error
}
