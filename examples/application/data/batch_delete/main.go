// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

func main() {
	app := application.NewApplication("c_c9c02f05c7c14131a6b4", "b85740891cbd419c86887c47082e9e13").Env(constants.PlatformEnvDEV)

	err := app.Data.Object("testObject").BatchDelete(context.Background(), []int64{1742643274486903, 1742643274487943})
	if err != nil {
		panic(err)
	}
}
