package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvUAT)

	result, err := app.Resources.File.DownloadAvatar(context.Background(), "debd656b74a54256959994b1c4d5482b_o")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %+v\n", len(result))
}
