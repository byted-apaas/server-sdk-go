// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package application

import (
	cHttp "github.com/byted-apaas/server-common-go/http"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/data"
	dataImpl "github.com/byted-apaas/server-sdk-go/service/data/impl"
	"github.com/byted-apaas/server-sdk-go/service/function"
	funcitonV2Impl "github.com/byted-apaas/server-sdk-go/service/function/impl"
	"github.com/byted-apaas/server-sdk-go/service/integration"
	"github.com/byted-apaas/server-sdk-go/service/resources_v2"
	"github.com/byted-apaas/server-sdk-go/service/tools"
)

type Application struct {
	Data      data.IDataV2
	Resources *resources_v2.Resources
	Tools     tools.ITools
	appCtx    *structs.AppCtx
	Integration integration.IIntegration
}

func NewApplication(clientID, clientSecret string) *Application {
	appCtx := &structs.AppCtx{
		Mode:       structs.AppModeOpenSDK,
		Credential: cHttp.NewAppCredential(clientID, clientSecret),
	}
	return &Application{
		Data:      dataImpl.NewDataV2(appCtx),
		Resources: resources_v2.NewResources(appCtx),
		Tools:     tools.NewTools(appCtx),
		Integration: integration.NewIntegration(appCtx),
		appCtx:    appCtx,
	}
}

func (a *Application) Function(apiName string) function.IFunction {
	return funcitonV2Impl.NewFunction(a.appCtx, apiName)
}

// Env 设置平台的环境，当访问平台的非正式环境时使用
func (a *Application) Env(env constants.PlatformEnvType) *Application {
	a.appCtx.Env = env
	return a
}
