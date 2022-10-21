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
	PageLimitMax = 200
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
