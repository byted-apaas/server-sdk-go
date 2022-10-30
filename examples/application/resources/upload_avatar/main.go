// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

func main() {
	app := application.NewApplication("xxx", "xxx").Env(constants.PlatformEnvPRE)

	result, err := app.Resources.File.UploadAvatarByPath(context.Background(), "avatar.png", "/Users/zhouweixin/Downloads/avatar.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", result)
}
