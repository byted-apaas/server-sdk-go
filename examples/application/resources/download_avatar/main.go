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
	app := application.NewApplication("c_e5447a833b4444969f58", "f05f672c28174735ad8fc28060b01b37").Env(constants.PlatformEnvUAT)

	result, err := app.Resources.File.DownloadAvatar(context.Background(), "debd656b74a54256959994b1c4d5482b_o")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", len(result))
}
