// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package global_config

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

func GetVar(ctx context.Context, key string) (string, error) {
	return GetVariable(ctx, nil, key)
}

func GetVariable(ctx context.Context, appCtx *structs.AppCtx, key string) (string, error) {
	return request.GetInstance(ctx).GetGlobalConfig(ctx, appCtx, key)
}
