// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package impl

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/function"
)

func Function(apiName string) function.IFunction {
	return NewFunction(nil, apiName)
}

type FunctionObject struct {
	appCtx  *structs.AppCtx
	apiName string
}

func NewFunction(s *structs.AppCtx, apiName string) function.IFunction {
	return &FunctionObject{appCtx: s, apiName: apiName}
}

func (f *FunctionObject) Invoke(ctx context.Context, params map[string]interface{}, result interface{}) error {
	return request.GetInstance(ctx).InvokeFunctionWithAuth(ctx, f.appCtx, f.apiName, params, result)
}
