package tenant

import (
	"context"

	cStructs "github.com/byted-apaas/server-common-go/structs"
)

type ITenant interface {
	// GetTenantInfo 获取租户信息
	GetTenantInfo(ctx context.Context) (*cStructs.Tenant, error)
}
