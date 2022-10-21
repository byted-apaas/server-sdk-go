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

	content, err := app.Resources.File.Download(context.Background(), "55edbb98a55e45faa92211dc40578970")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(content))
}
