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

	result, err := app.Resources.File.DownloadAvatar(context.Background(), "debd656b74a54256959994b1c4d5482b_o")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", len(result))
}
