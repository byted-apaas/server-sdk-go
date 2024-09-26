package main

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	ctx := context.Background()
	// 在这里补充业务代码
	app := application.NewApplication("***", "***").Env(constants.PlatformEnvLR)
	ctx = utils.LocalDebugMode(ctx)

	var result interface{}
	err := app.Function("returnBoolFun").Invoke(ctx, map[string]interface{}{"inputText": "小白"}, &result)
	if err != nil {
		panic(err)
	}
}
