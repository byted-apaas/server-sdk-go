// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package data

import (
	"reflect"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data"
	"github.com/google/uuid"
)

type TransactionObject struct {
	transaction   *Transaction
	objectAPIName string
}

func NewTransactionObject(t *Transaction, objectAPIName string) data.ITransactionObject {
	return &TransactionObject{transaction: t, objectAPIName: objectAPIName}
}

func (t *TransactionObject) RegisterCreate(record interface{}) (*structs.TransactionRecordID, error) {
	if err := t.check(); err != nil {
		return nil, err
	}

	if record == nil {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterCreate] The record cannot be empty")
		return nil, t.transaction.err
	}

	input := map[string]interface{}{}
	err := cUtils.Decode(record, &input)
	if err != nil {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterCreate] Decode record failed, err: %+v", err)
		return nil, t.transaction.err
	}

	random, err := uuid.NewRandom()
	if err != nil {
		t.transaction.err = cExceptions.InternalError("[RegisterCreate] The create uuid failed, err: %+v", err)
		return nil, t.transaction.err
	}

	uuidValue := random.String()
	input["_id"] = uuidValue

	operation, err := newTransactionOperation(constants.OperationTypeCreate, t.objectAPIName, input)
	if err != nil {
		t.transaction.err = err
		return nil, t.transaction.err
	}

	result := &structs.TransactionRecordID{ID: uuidValue}
	t.transaction.placeholders[uuidValue] = 0
	t.transaction.uuidToResult[uuidValue] = result
	t.transaction.operations = append(t.transaction.operations, operation)
	return result, nil
}

func (t *TransactionObject) RegisterUpdate(_id interface{}, record interface{}) {
	if err := t.check(); err != nil {
		return
	}

	input := map[string]interface{}{}
	err := cUtils.Decode(record, &input)
	if err != nil {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterUpdate] Decode record failed, err: %+v", err)
		return
	}

	input["_id"] = _id
	operation, err := newTransactionOperation(constants.OperationTypeUpdate, t.objectAPIName, input)
	if err != nil {
		t.transaction.err = err
		return
	}

	t.transaction.operations = append(t.transaction.operations, operation)
}

func (t *TransactionObject) RegisterDelete(_id interface{}) {
	if err := t.check(); err != nil {
		return
	}

	operation, err := newTransactionOperation(constants.OperationTypeDelete, t.objectAPIName, map[string]interface{}{"_id": _id})
	if err != nil {
		t.transaction.err = err
		return
	}
	t.transaction.operations = append(t.transaction.operations, operation)
}

func (t *TransactionObject) RegisterBatchCreate(records interface{}) ([]interface{}, error) {
	if err := t.check(); err != nil {
		return nil, err
	}

	var newRecords []map[string]interface{}
	if err := cUtils.Decode(records, &newRecords); err != nil {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterBatchCreate] The type of records is not []map[string]interface{}")
		return nil, t.transaction.err
	}

	result := make([]interface{}, len(newRecords))
	for i, record := range newRecords {
		if record == nil {
			continue
		}

		random, err := uuid.NewRandom()
		if err != nil {
			t.transaction.err = cExceptions.InternalError("[RegisterBatchCreate] The create uuid failed, err: %+v", err)
			return nil, t.transaction.err
		}

		uuidValue := random.String()
		record["_id"] = uuidValue
		result[i] = uuidValue
		t.transaction.placeholders[uuidValue] = 0
	}
	t.transaction.batchResults = append(t.transaction.batchResults, result)

	operation, err := newTransactionOperation(constants.OperationTypeBatchCreate, t.objectAPIName, newRecords)
	if err != nil {
		t.transaction.err = err
		return nil, t.transaction.err
	}
	t.transaction.operations = append(t.transaction.operations, operation)
	return result, nil
}

func (t *TransactionObject) RegisterBatchUpdate(records interface{}) {
	if err := t.check(); err != nil {
		return
	}

	var newRecords []map[string]interface{}
	if err := cUtils.Decode(records, &newRecords); err != nil {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterBatchUpdate] The type of records is not []map[string]interface{}")
		return
	}

	if len(newRecords) == 0 {
		return
	}

	isSameValue := true
	var _ids []interface{}
	flagRecord := newRecords[0]
	flagRecordID, ok := flagRecord["_id"]
	if !ok {
		t.transaction.err = cExceptions.InvalidParamError("[RegisterBatchUpdate] The record._id is empty")
		return
	}
	delete(flagRecord, "_id")
	_ids = append(_ids, flagRecordID)
	for i := 1; i < len(newRecords); i++ {
		_id, ok := newRecords[i]["_id"]
		if !ok {
			t.transaction.err = cExceptions.InvalidParamError("[RegisterBatchUpdate] The record._id is empty")
			return
		}
		delete(newRecords[i], "_id")
		if !reflect.DeepEqual(flagRecord, newRecords) {
			isSameValue = false
			newRecords[i]["_id"] = _id
			break
		}
		newRecords[i]["_id"] = _id
		_ids = append(_ids, _id)
	}
	flagRecord["_id"] = flagRecordID

	op, input := constants.OperationType(constants.OperationTypeBatchUpdate), interface{}(newRecords)
	if isSameValue {
		op, input = constants.OperationTypeBatchUpdateSameValue, map[string]interface{}{"ids": _ids, "value": flagRecord}
	}

	operation, err := newTransactionOperation(op, t.objectAPIName, input)
	if err != nil {
		t.transaction.err = err
		return
	}
	t.transaction.operations = append(t.transaction.operations, operation)
}

func (t *TransactionObject) RegisterBatchDelete(_ids interface{}) {
	if err := t.check(); err != nil {
		return
	}

	var (
		operation *structs.TransactionOperation
		err       error
	)

	switch _ids.(type) {
	case []interface{}:
		ids, _ := _ids.([]interface{})
		if len(ids) == 0 {
			return
		}
		operation, err = newTransactionOperation(constants.OperationTypeBatchDelete, t.objectAPIName, map[string]interface{}{"ids": _ids})
	case []int64:
		ids, _ := _ids.([]int64)
		if len(ids) == 0 {
			return
		}
		operation, err = newTransactionOperation(constants.OperationTypeBatchDelete, t.objectAPIName, map[string]interface{}{"ids": _ids})
	case []string:
		ids, _ := _ids.([]string)
		if len(ids) == 0 {
			return
		}
		operation, err = newTransactionOperation(constants.OperationTypeBatchDelete, t.objectAPIName, map[string]interface{}{"ids": _ids})
	default:
		t.transaction.err = cExceptions.InvalidParamError("[RegisterBatchDelete] The type of _ids should be []interface{}, []int64 or []string, but %s", reflect.TypeOf(_ids))
		return
	}

	if err != nil {
		t.transaction.err = err
		return
	}
	t.transaction.operations = append(t.transaction.operations, operation)
}

func (t *TransactionObject) check() error {
	if t.transaction.err != nil {
		return t.transaction.err
	}

	if t.transaction.isCommitted {
		t.transaction.err = cExceptions.InvalidParamError("[check] Committed transaction cannot register operations")
		return t.transaction.err
	}

	return nil
}
