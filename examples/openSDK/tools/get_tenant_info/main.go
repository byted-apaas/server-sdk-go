package main

import (
	"context"
	"fmt"

	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/opensdk"
)

func main() {
	app := application.NewApplication("***", "***").Env(constants.PlatformEnvUAT)

	tenant, err := app.Tools.GetTenantInfo(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("tenant: %+v\n", tenant)
}
