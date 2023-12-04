package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	app := opensdk.NewApplication("***", "***").Env(constants.PlatformEnvUAT)

	content, err := app.Resources.File.Download(context.Background(), "55edbb98a55e45faa92211dc40578970")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(content))
}
