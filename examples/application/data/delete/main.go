// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

func main() {
	app := application.NewApplication("xxx", "xxx").Env(constants.PlatformEnvPRE)

	err := app.Data.Object("testObject").Delete(context.Background(), 1742602778895479)
	if err != nil {
		panic(err)
	}
}
