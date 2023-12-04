// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package constants

import (
	"fmt"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

type PlatformEnvType int

const (
	PlatformEnvDEV PlatformEnvType = iota + 1
	PlatformEnvUAT
	PlatformEnvLR
	PlatformEnvPRE
	PlatformEnvOnline
)

func (p PlatformEnvType) String() string {
	switch p {
	case PlatformEnvUAT:
		return cConstants.EnvTypeStaging
	case PlatformEnvLR:
		return cConstants.EnvTypeLr
	case PlatformEnvPRE:
		return cConstants.EnvTypeGray
	case PlatformEnvOnline:
		return cConstants.EnvTypeOnline
	}
	panic(fmt.Sprintf("invalid platform env type %d", p))
}

const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

const (
	// PageLimitDefault SDK 仅作 limit 小于 0 校验，其他校验拦截由底层控制，如当查询结果不满足受控条件时，limit 仅支持最大 200
	PageLimitDefault = 200
)

type OperationType int

const (
	OperationTypeCreate = iota + 1
	OperationTypeUpdate
	OperationTypeDelete
	OperationTypeBatchCreate
	OperationTypeBatchUpdate
	OperationTypeBatchDelete
	OperationTypeBatchUpdateSameValue
)

const (
	GlobalVariableCacheTableKey = "global-variable-cache-table"
)
