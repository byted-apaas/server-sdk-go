// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"context"
	"reflect"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/data"
)

type Transaction struct {
	appCtx       *structs.AppCtx                         // 应用级凭证
	placeholders map[string]int64                        // 占位符：uuid -> id
	uuidToResult map[string]*structs.TransactionRecordID // 单条注册创建的结果
	batchResults [][]interface{}                         // 批量注册创建的结果
	operations   []*structs.TransactionOperation         // 操作集
	isCommitted  bool                                    // 是否已提交
	err          error                                   // 错误
}

func NewTransaction(s *structs.AppCtx) data.ITransaction {
	t := Transaction{appCtx: s}
	t.placeholders = make(map[string]int64)
	t.uuidToResult = make(map[string]*structs.TransactionRecordID)
	t.isCommitted = false
	return &t
}

func (t *Transaction) Object(objectAPIName string) data.ITransactionObject {
	return NewTransactionObject(t, objectAPIName)
}

func (t *Transaction) Commit(ctx context.Context) error {
	if t.err != nil {
		return t.err
	}

	if t.isCommitted {
		return cExceptions.InvalidParamError("The transaction cannot be committed repeatedly")
	}
	t.isCommitted = true

	if len(t.operations) == 0 {
		return nil
	}

	uuidToRecordID, err := request.GetInstance(ctx).Transaction(ctx, t.appCtx, t.placeholders, t.operations)
	if err != nil {
		return err
	}

	// 单个
	for uuid := range t.uuidToResult {
		recordID, ok := uuidToRecordID[uuid]
		if !ok {
			return cExceptions.InternalError(`The uuid (%s) is not in uuidToRecordID (%+v)`, uuid, uuidToRecordID)
		}
		t.uuidToResult[uuid].ID = recordID
	}

	// 批量
	for i, uuids := range t.batchResults {
		for j, rawUUID := range uuids {
			uuid, ok := rawUUID.(string)
			if !ok {
				return cExceptions.InternalError(`The type of uuid should be string, but %v`, reflect.TypeOf(uuid))
			}

			recordID, ok := uuidToRecordID[uuid]
			if !ok {
				return cExceptions.InternalError(`The uuid (%s) is not in uuidToRecordID (%+v)`, uuid, uuidToRecordID)
			}
			t.batchResults[i][j] = recordID
		}
	}

	return nil
}
