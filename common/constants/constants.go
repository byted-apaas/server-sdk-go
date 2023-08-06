// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package constants

import (
	"fmt"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

type PlatformEnvType int

const (
	PlatformEnvLR PlatformEnvType = iota + 1
	PlatformEnvPRE
	PlatformEnvOnline
)

func (p PlatformEnvType) String() string {
	switch p {
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
	OpRateLimitError = "k_op_ec_20003"
	FsRateLimitError = "k_fs_ec_000004"
)

const (
	GlobalVariableCacheTableKey = "global-variable-cache-table"
)

type ProcessAuthFieldType int64

const (
	ProcessAuthFieldType_Default     ProcessAuthFieldType = 0
	ProcessAuthFieldType_BothResult  ProcessAuthFieldType = 1
	ProcessAuthFieldType_SliceResult ProcessAuthFieldType = 2
	ProcessAuthFieldType_MapResult   ProcessAuthFieldType = 3
)

type SetSystemMod int64

const (
	SetSystemMod_NotSet  SetSystemMod = 0
	SetSystemMod_FullSet SetSystemMod = 1
	SetSystemMod_Other   SetSystemMod = 2
)

type CommitSetSystemMod int64

const (
	CommitSetSystemMod_NotSet      CommitSetSystemMod = 0
	CommitSetSystemMod_SysFieldSet CommitSetSystemMod = 1
)
