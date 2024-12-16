package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvUAT)

	result, err := app.Resources.File.UploadAvatarByPath(context.Background(), "avatar.png", "/Users/zhouweixin/Downloads/avatar.png")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", result)
}
