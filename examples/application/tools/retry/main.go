// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"time"

	"github.com/byted-apaas/server-sdk-go/application"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/tools"
)

func main() {
	app := application.NewApplication("c_e5447a833b4444969f58", "f05f672c28174735ad8fc28060b01b37").Env(constants.PlatformEnvUAT)

	executeTimes := 0
	// 失败时需要重试的函数
	doSomething := func() error {
		executeTimes += 1
		if executeTimes <= 3 {
			fmt.Printf("执行次数：%d，出错了\n", executeTimes)
			return fmt.Errorf("故意出错，测试重试工具")
		}

		fmt.Printf("执行次数：%d，成功了\n", executeTimes)
		return nil
	}

	if err := app.Tools.Retry(doSomething, &tools.RetryOption{
		RetryCount:    5,                      // 重试次数
		RetryInterval: 500 * time.Millisecond, // 重试间隔
	}); err != nil {
		panic(err)
	}
}
