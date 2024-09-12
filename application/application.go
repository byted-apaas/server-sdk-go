// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package application

import (
	"github.com/byted-apaas/server-common-go/logger"
	appImpl "github.com/byted-apaas/server-sdk-go/service/app/impl"
	dataImpl "github.com/byted-apaas/server-sdk-go/service/data/impl"
	eventImpl "github.com/byted-apaas/server-sdk-go/service/event/impl"
	"github.com/byted-apaas/server-sdk-go/service/flow"
	globalVarImpl "github.com/byted-apaas/server-sdk-go/service/global_var/impl"
	"github.com/byted-apaas/server-sdk-go/service/integration"
	messageImpl "github.com/byted-apaas/server-sdk-go/service/message/impl"
	"github.com/byted-apaas/server-sdk-go/service/metadata"
	"github.com/byted-apaas/server-sdk-go/service/resources"
	tenantImpl "github.com/byted-apaas/server-sdk-go/service/tenant/impl"
	"github.com/byted-apaas/server-sdk-go/service/user"
)

var (
	Tenant = tenantImpl.NewTenant(nil)
	App    = appImpl.NewApp(nil)
	User   = user.NewUser(nil)
	Event  = eventImpl.NewEvent(nil)

	Data      = dataImpl.NewData(nil)
	DataV3    = dataImpl.NewDataV3(nil)
	Metadata  = metadata.NewMetadata(nil)
	GlobalVar = globalVarImpl.NewGlobalVar(nil)
	Flow      = flow.NewFlow(nil)
	Resources = resources.NewResources(nil)
	Msg       = messageImpl.NewMsg(nil)
	GetLogger = logger.GetLogger
	// GetContext 获得一些上游传入的系统上下文参数
	GetContext    = user.GetContext
	GetContextMap = user.GetContextMap
	Integration   = integration.NewIntegration(nil)
)
