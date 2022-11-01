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

type ITools interface {
	// Retry 重试工具, 默认一次立即重试 (option is nil)
	Retry(f func() error, option *RetryOption) error
	// GetTenantInfo 获取租户信息工具
	GetTenantInfo(ctx context.Context) (*cStructs.Tenant, error)
	// SetLaneNameToCtx 设置泳道环境, 只在调试模式下生效
	SetLaneNameToCtx(ctx context.Context, lane string) context.Context

	// ParseErr 转换 err 为有结构的错误，若 err 非系统返回的，则 code 为 ErrCodeDeveloperError
	ParseErr(ctx context.Context, err error) *exceptions.BaseError
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
