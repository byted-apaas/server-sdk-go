// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import (
	cHttp "github.com/byted-apaas/server-common-go/http"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

type AppCtx struct {
	Mode        AppMode
	Env         constants.PlatformEnvType
	Credential  *cHttp.AppCredential
	DataVersion DataVersion // data 版本
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

func (a *AppCtx) IsDataV3() bool {
	return a.DataVersion == DataVersionV3
}

type AppMode int

const (
	AppModeOpenSDK AppMode = iota + 1
	AppModeFaaSSDK
)

type DataVersion string

const (
	DataVersionV1 = "v1"
	DataVersionV2 = "v2"
	DataVersionV3 = "v3"
)
