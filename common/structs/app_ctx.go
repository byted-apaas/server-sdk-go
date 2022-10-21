// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import (
	cHttp "github.com/byted-apaas/server-common-go/http"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

type AppCtx struct {
	Mode       AppMode
	Env        constants.PlatformEnvType
	Credential *cHttp.AppCredential
}

func (a *AppCtx) IsOpenSDK() bool {
	if a != nil && a.Mode == AppModeOpenSDK {
		return true
	}
	return false
}

func (a *AppCtx) GetEnv() constants.PlatformEnvType {
	if a != nil && a.Env != 0 {
		return a.Env
	}
	return constants.PlatformEnvOnline
}

type AppMode int

const (
	AppModeOpenSDK AppMode = iota + 1
	AppModeFaaSSDK
)
