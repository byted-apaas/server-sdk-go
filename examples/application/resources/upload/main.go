// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
)

func main() {
	app := application.NewApplication("c_e5447a833b4444969f58", "f05f672c28174735ad8fc28060b01b37").Env(constants.PlatformEnvUAT)

	result1, err := app.Resources.File.UploadByPath(context.Background(), "main.go", "./main.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", result1)

	reader := bytes.NewReader([]byte("测试 reader"))
	result2, err := app.Resources.File.UploadByReader(context.Background(), "main.go", reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", result2)

	data, err := ioutil.ReadAll(bytes.NewReader([]byte("测试 buffer")))
	if err != nil {
		panic(err)
	}
	result3, err := app.Resources.File.UploadByBuffer(context.Background(), "main.go", data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result3: %+v\n", result3)
}
