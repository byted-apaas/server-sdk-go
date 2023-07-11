package impl

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type App struct {
	appCtx *structs.AppCtx
}

func NewApp(appCtx *structs.AppCtx) *App {
	return &App{appCtx: appCtx}
}

func (a *App) GetAppInfo(ctx context.Context) (*cStructs.AppInfo, error) {
	return cUtils.GetAppInfoFromCtx(ctx)
}
