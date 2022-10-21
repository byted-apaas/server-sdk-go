// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

func main() {
	app := application.NewApplication("c_a81a2dbc72f447478ead", "b539814d85bb48efb54b9bc6201c7160").Env(constants.PlatformEnvDEV)
	fmt.Println(app)
}
