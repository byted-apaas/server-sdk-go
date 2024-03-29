// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package api

import (
	"github.com/byted-apaas/server-common-go/logger"
	dataImpl "github.com/byted-apaas/server-sdk-go/service/data/impl"
	"github.com/byted-apaas/server-sdk-go/service/flow"
	functionV2Impl "github.com/byted-apaas/server-sdk-go/service/function/impl"
	"github.com/byted-apaas/server-sdk-go/service/global_config"
	"github.com/byted-apaas/server-sdk-go/service/integration"
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
	Integration = integration.NewIntegration(nil)
	//GetContext 获得一些上游传入的系统上下文参数
	GetContext    = user.GetContext
	GetContextMap = user.GetContextMap
)
