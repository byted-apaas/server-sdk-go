package impl

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/service/global_config"
	"github.com/byted-apaas/server-sdk-go/service/global_var"
)

type GlobalVar struct {
	appCtx *structs.AppCtx
}

func NewGlobalVar(s *structs.AppCtx) global_var.IGlobalVar {
	return &GlobalVar{appCtx: s}
}

func (GlobalVar) GetVar(ctx context.Context, key string) (string, error) {
	return global_config.GetVar(ctx, key)
}

func (g *GlobalVar) GetVariable(ctx context.Context, key string) (string, error) {
	return global_config.GetVariable(ctx, g.appCtx, key)
}
