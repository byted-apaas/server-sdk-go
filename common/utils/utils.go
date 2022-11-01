// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package utils

import (
	"context"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

// LocalDebugMode 设置调试模式
func LocalDebugMode(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return cUtils.SetDebugTypeToCtx(ctx, cConstants.DebugTypeLocal)
}

// SetLaneNameToCtx 设置泳道环境, 只在调试模式下生效
func SetLaneNameToCtx(ctx context.Context, lane string) context.Context {
	if !cUtils.IsDebug(ctx) {
		return ctx
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return cUtils.SetTTEnvToCtx(ctx, lane)
}
