// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvDEV)

	err := app.Data.Object("testObject").BatchDelete(context.Background(), []int64{1742643274486903, 1742643274487943})
	if err != nil {
		panic(err)
	}
}
