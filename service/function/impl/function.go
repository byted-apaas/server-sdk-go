// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package impl

import (
	"context"

	"github.com/byted-apaas/server-common-go/logger"
	cUtils "github.com/byted-apaas/server-common-go/utils"
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

func PrintInvokeReqInfo(ctx context.Context, f *FunctionObject, params map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			logger.GetLogger(ctx).Errorf("PrintInvokeReqInfo panic recover: %+v", r)
		}
	}()
	if err != nil {
		// print log for debug
		// 1. print err、apiName、reqParams、mode、credentialID
		var innerErr error
		var mode structs.AppMode
		var credentialID string
		if f.appCtx != nil {
			mode = f.appCtx.Mode
			logger.GetLogger(ctx).Infof("PrintInvokeReqInfo f.appCtx.Mode: %v", mode)
		}
		if f.appCtx != nil && f.appCtx.Credential != nil {
			credentialID = f.appCtx.Credential.GetID()
			logger.GetLogger(ctx).Infof("PrintInvokeReqInfo f.appCtx.Credential id: %v", credentialID)
		}
		logger.GetLogger(ctx).Errorf("InvokeFunctionWithAuth failed, err: %v; apiName: %v, params: %+v, appCtx.mode: %v, appCtx.Credential.id: %v", err, f.apiName, params, mode, credentialID)

		// 2. print more info: try get TenantInfo and userInfo and print them
		var tenant *structs.TenantInfo
		var namespace string
		if f.appCtx != nil && f.appCtx.Credential != nil {
			tenant, innerErr = f.appCtx.Credential.GetTenantInfo(ctx)
			if innerErr != nil {
				// 这个错误是在主体流程错误时，额外读取参数时的报错，不要作为 debug 的错误
				logger.GetLogger(ctx).Infof("PrintInvokeReqInfo f.appCtx.Credential.GetTenantInfo failed, err: %v", innerErr)
			}
			if tenant != nil {
				namespace = tenant.Namespace
			}
		} else {
			namespace = cUtils.GetNamespace()
		}
		if tenant == nil {
			logger.GetLogger(ctx).Infof("PrintInvokeReqInfo tenant is nil, set empty tenant info just for log")
			tenant = &structs.TenantInfo{}
		}
		userID := cUtils.GetUserIDFromCtx(ctx)
		logger.GetLogger(ctx).Errorf("InvokeFunctionWithAuth failed, tenant: %v, namespace: %v, userID: %v", *tenant, namespace, userID)
	}

}

func (f *FunctionObject) Invoke(ctx context.Context, params map[string]interface{}, result interface{}) error {
	err := request.GetInstance(ctx).InvokeFunctionWithAuth(ctx, f.appCtx, f.apiName, params, result)
	PrintInvokeReqInfo(ctx, f, params, err)
	return err
}
