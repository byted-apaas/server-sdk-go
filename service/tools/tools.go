// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package tools

import (
	"context"
	"time"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cStructs "github.com/byted-apaas/server-common-go/structs"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/request"
)

const (
	CtxKeyTenantID   = "KTenantID"
	CtxKeyTenantType = "KTenantType"
	CtxKeyUserID     = "KUserID"
	CtxKeyNamespace  = "KNamespace"
)

type ITools interface {
	// Retry 重试工具, 默认一次立即重试 (option is nil)
	Retry(f func() error, option *RetryOption) error
	// GetTenantInfo 获取租户信息工具
	GetTenantInfo(ctx context.Context) (*cStructs.Tenant, error)
	// SetLaneNameToCtx 设置泳道环境, 只在调试模式下生效
	SetLaneNameToCtx(ctx context.Context, lane string) context.Context

	// ParseErr 转换 err 为有结构的错误，若 err 非系统返回的，则 code 为 ErrCodeDeveloperError
	ParseErr(ctx context.Context, err error) *exceptions.BaseError

	// GetCommonReqInfo 获取通用的请求信息，包括租户id，租户类型，namesapce, userID
	GetCommonReqInfo(ctx context.Context) (tenantID int64, tenantType int64, namespace string, userID int64)
}

// RetryOption 重试选项
type RetryOption struct {
	RetryCount    int           // 重试次数
	RetryInterval time.Duration // 重试间隔
}

func NewRetryOption(retryCount int, retryDelay time.Duration) *RetryOption {
	return &RetryOption{RetryCount: retryCount, RetryInterval: retryDelay}
}

type Tools struct {
	appCtx *structs.AppCtx
}

func NewTools(s *structs.AppCtx) ITools {
	return &Tools{appCtx: s}
}

func (t *Tools) Retry(f func() error, options *RetryOption) error {
	if options == nil {
		options = NewRetryOption(1, 0)
	}
	return cUtils.InvokeFuncWithRetry(options.RetryCount, options.RetryInterval, f)
}

func (t *Tools) GetTenantInfo(ctx context.Context) (*cStructs.Tenant, error) {
	if t.appCtx == nil {
		t.appCtx = &structs.AppCtx{
			Mode: structs.AppModeFaaSSDK,
		}
	}
	return request.GetInstance(ctx).GetTenantInfo(ctx, t.appCtx)
}

func (t *Tools) SetLaneNameToCtx(ctx context.Context, lane string) context.Context {
	return utils.SetLaneNameToCtx(ctx, lane)
}

func (t *Tools) ParseErr(ctx context.Context, err error) *exceptions.BaseError {
	return cExceptions.ParseErrForUser(err)
}

func (t *Tools) GetCommonReqInfo(ctx context.Context) (tenantID int64, tenantType int64, namespace string, userID int64) {
	if tID, ok := ctx.Value(CtxKeyTenantID).(int64); ok {
		tenantID = tID
	}

	if tType, ok := ctx.Value(CtxKeyTenantType).(int64); ok {
		tenantType = tType
	}

	if ns, ok := ctx.Value(CtxKeyNamespace).(string); ok {
		namespace = ns
	}

	if uID, ok := ctx.Value(CtxKeyUserID).(int64); ok {
		userID = uID
	}
	return
}
