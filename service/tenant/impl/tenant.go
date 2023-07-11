package impl

import (
	"context"

	cStructs "github.com/byted-apaas/server-common-go/structs"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/tenant"
)

type Tenant struct {
	appCtx *structs.AppCtx
}

func NewTenant(s *structs.AppCtx) tenant.ITenant {
	return &Tenant{appCtx: s}
}

func (t Tenant) GetTenantInfo(ctx context.Context) (*cStructs.Tenant, error) {
	if t.appCtx == nil {
		t.appCtx = &structs.AppCtx{
			Mode: structs.AppModeFaaSSDK,
		}
	}
	return request.GetInstance(ctx).GetTenantInfo(ctx, t.appCtx)
}
