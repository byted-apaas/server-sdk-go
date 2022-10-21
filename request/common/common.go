// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package common

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
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
		"triggerType": "public-api",
	}, nil
}

func BuildInvokeParamsObj(ctx context.Context, apiName string, params interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"apiAlias":    apiName,
		"params":      params,
		"context":     BuildInvokeSysParams(ctx),
		"triggerType": "public-api",
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
