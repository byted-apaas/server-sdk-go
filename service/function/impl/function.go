// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package impl

import (
	"context"

	cHttp "github.com/byted-apaas/server-common-go/http"
	"github.com/byted-apaas/server-common-go/logger"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/function"
	"github.com/byted-apaas/server-sdk-go/service/tools"
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
	err := request.GetInstance(ctx).InvokeFunctionWithAuth(ctx, f.appCtx, f.apiName, params, result)
	if err != nil {
		var mode structs.AppMode
		credential := &cHttp.AppCredential{}
		if f.appCtx != nil {
			mode = f.appCtx.Mode
		}
		if f.appCtx != nil && f.appCtx.Credential != nil {
			credential = f.appCtx.Credential
		}
		tools := tools.NewTools(f.appCtx)
		tenantID, tenantType, namespace, userID := tools.GetCommonReqInfo(ctx)
		logger.GetLogger(ctx).Errorf("InvokeFunctionWithAuth failed, reqInfo: tenantID: %v, tenantType: %v, namespace: %v, userID: %v", tenantID, tenantType, namespace, userID)
		logger.GetLogger(ctx).Errorf("InvokeFunctionWithAuth failed, err: %v; apiName: %v, params: %+v, appCtx.mode: %v, appCtx.Credential.id: %v", err, f.apiName, params, mode, credential.GetID())
		return err
	}
	return nil
}
