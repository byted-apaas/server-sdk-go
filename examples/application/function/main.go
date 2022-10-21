// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/utils"
)

func main() {
	ctx := context.Background()
	// 在这里补充业务代码
	app := application.NewApplication("c_d2d7429d1ddb4755a973", "a19303a91c224d80a765f54c9ac86abb").Env(constants.PlatformEnvLR)
	ctx = utils.LocalDebugMode(ctx)

	var result interface{}
	err := app.Function("returnBoolFun").Invoke(ctx, map[string]interface{}{"inputText": "小白"}, &result)
	if err != nil {
		panic(err)
	}
}
