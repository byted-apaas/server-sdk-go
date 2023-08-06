package impl

import (
	"context"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type App struct {
	appCtx *structs.AppCtx
}

func NewApp(appCtx *structs.AppCtx) *App {
	return &App{appCtx: appCtx}
}

func (a *App) GetAppInfo(ctx context.Context) (*structs.AppInfo, error) {
	return cUtils.GetAppInfoFromCtx(ctx)
}
