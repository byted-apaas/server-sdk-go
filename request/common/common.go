// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package common

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
)

func BuildInvokeParamsStr(ctx context.Context, apiName string, params interface{}) (map[string]interface{}, error) {
	sysParams, bizParams, err := BuildInvokeParamAndContext(ctx, params)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"apiAlias":    apiName,
		"params":      bizParams,
		"context":     sysParams,
		"triggerType": "default-trigger",
	}, nil
}

func BuildInvokeParamsObj(ctx context.Context, apiName string, params interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"apiAlias":    apiName,
		"params":      params,
		"context":     BuildInvokeSysParams(ctx),
		"triggerType": "default-trigger",
	}, nil
}

func BuildInvokeParamAndContext(ctx context.Context, params interface{}) (string, string, error) {
	sysParams, _ := cUtils.JsonMarshalBytes(BuildInvokeSysParams(ctx))

	bizParams, err := cUtils.JsonMarshalBytes(params)
	if err != nil {
		return "", "", cExceptions.InvalidParamError("Marshal params failed, err: %+v", err)
	}
	return string(sysParams), string(bizParams), nil
}

func BuildInvokeSysParams(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"triggertaskid":             cUtils.GetTriggerTaskIDFromCtx(ctx),
		"x-kunlun-distributed-mask": cUtils.GetDistributedMaskFromCtx(ctx),
		"x-kunlun-loop-masks":       cUtils.GetLoopMaskFromCtx(ctx),
	}
}

func GenBatchResultByRecords(records map[int64]interface{}, errMap map[int64]string) *structs.BatchResult {
	result := &structs.BatchResult{Code: "", Msg: "success"}
	for id := range records {
		result.Data = append(result.Data, GenBatchResultData(id, errMap))
	}
	return result
}

func GenBatchResultByRecordIDs(recordIDs []int64, errMap map[int64]string) *structs.BatchResult {
	result := &structs.BatchResult{Code: "", Msg: "success"}
	for _, id := range recordIDs {
		result.Data = append(result.Data, GenBatchResultData(id, errMap))
	}
	return result
}

func GenBatchResultData(id int64, errMap map[int64]string) structs.BatchResultData {
	d := structs.BatchResultData{ID: id, Success: true}
	if errCode, ok := errMap[id]; ok {
		d.Success = false
		d.Errors = []structs.BatchResultDataError{
			{
				Code: errCode,
			},
		}
	}
	return d
}

func GenBatchResultByRecordsV3(records map[string]interface{}, errMap map[string]string) *structs.BatchResultV3 {
	result := &structs.BatchResultV3{Code: "", Msg: "success"}
	for id := range records {
		result.Data = append(result.Data, GenBatchResultDataV3(id, errMap))
	}
	return result
}

func GenBatchResultByRecordIDsV3(recordIDs []string, errMap map[string]string) *structs.BatchResultV3 {
	result := &structs.BatchResultV3{Code: "", Msg: "success"}
	for _, id := range recordIDs {
		result.Data = append(result.Data, GenBatchResultDataV3(id, errMap))
	}
	return result
}

func GenBatchResultDataV3(id string, errMap map[string]string) structs.BatchResultDataV3 {
	d := structs.BatchResultDataV3{ID: id, Success: true}
	if errCode, ok := errMap[id]; ok {
		d.Success = false
		d.Errors = []structs.BatchResultDataError{
			{
				Code: errCode,
			},
		}
	}
	return d
}
