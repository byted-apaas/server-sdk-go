// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

// Package api is deprecated. Please use the new package instead.
//
// Deprecated: This package is no longer maintained. Please refer to the example below to use the other new packages.
//
// Example:
//
//  1. api(code.byted.org/byted-apaas/server-sdk-go) -> application(code.byted.org/byted-apaas/server-sdk-go)
//     api.Data -> application.Data
//     api.Metadata -> application.Metadata
//     api.Resources -> application.Resources
//     api.Msg -> application.Msg
//     api.GetLogger -> application.GetLogger
//     api.GetVar -> application.GlobalVar
//     api.Tools.GetTenantInfo -> application.Tenant.GetTenantInfo
//     api.User -> application.User
//     api.Flow -> application.Flow
//     api.GetContext -> application.GetContext
//     api.GetContextMap -> application.GetContextMap
//
//  2. api(code.byted.org/byted-apaas/server-sdk-go) -> baas(code.byted.org/byted-apaas/baas-sdk-go)
//     api.Tasks -> baas.Tasks
//
//  3. api(code.byted.org/byted-apaas/server-sdk-go) -> faas(code.byted.org/byted-apaas/faas-sdk-go)
//     api.Function -> faas.Function
//     api.Tools -> faas.Tools
package api

import (
	"github.com/byted-apaas/server-common-go/logger"
	dataImpl "github.com/byted-apaas/server-sdk-go/service/data/impl"
	"github.com/byted-apaas/server-sdk-go/service/flow"
	functionV2Impl "github.com/byted-apaas/server-sdk-go/service/function/impl"
	"github.com/byted-apaas/server-sdk-go/service/global_config"
	messageImpl "github.com/byted-apaas/server-sdk-go/service/message/impl"
	"github.com/byted-apaas/server-sdk-go/service/metadata"
	"github.com/byted-apaas/server-sdk-go/service/resources"
	tasksImpl "github.com/byted-apaas/server-sdk-go/service/tasks/impl"
	"github.com/byted-apaas/server-sdk-go/service/tools"
	"github.com/byted-apaas/server-sdk-go/service/user"
)

var (
	Data      = dataImpl.NewData(nil)
	Metadata  = metadata.NewMetadata(nil)
	Resources = resources.NewResources(nil)
	Tasks     = tasksImpl.NewTasks(nil)
	Msg       = messageImpl.NewMsg(nil)
	GetLogger = logger.GetLogger
	Function  = functionV2Impl.Function
	GetVar    = global_config.GetVar
	Tools     = tools.NewTools(nil)
	User      = user.NewUser(nil)
	Flow      = flow.NewFlow(nil)
	//GetContext 获得一些上游传入的系统上下文参数
	GetContext    = user.GetContext
	GetContextMap = user.GetContextMap
)
