package integration

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

type IIntegration interface {
	GetTenantAccessToken(ctx context.Context, apiName string) (*structs.TenantAccessToken, error)
	GetAppAccessToken(ctx context.Context, apiName string) (*structs.AppAccessToken, error)
	GetDefaultTenantAccessToken(ctx context.Context) (*structs.TenantAccessToken, error)
	GetDefaultAppAccessToken(ctx context.Context) (*structs.AppAccessToken, error)
}

type Integration struct {
	appCtx *structs.AppCtx
}

func (i Integration) GetTenantAccessToken(ctx context.Context, apiName string) (*structs.TenantAccessToken, error) {
	if apiName == "" {
		return nil, cExceptions.InvalidParamError("The apiName can not be empty")
	}
	return request.GetInstance(ctx).GetTenantAccessToken(ctx, i.appCtx, apiName)
}

func (i Integration) GetAppAccessToken(ctx context.Context, apiName string) (*structs.AppAccessToken, error) {
	if apiName == "" {
		return nil, cExceptions.InvalidParamError("The apiName can not be empty")
	}
	return request.GetInstance(ctx).GetAppAccessToken(ctx, i.appCtx, apiName)
}

func (i Integration) GetDefaultTenantAccessToken(ctx context.Context) (*structs.TenantAccessToken, error) {
	return request.GetInstance(ctx).GetDefaultTenantAccessToken(ctx, i.appCtx)
}

func (i Integration) GetDefaultAppAccessToken(ctx context.Context) (*structs.AppAccessToken, error) {
	return request.GetInstance(ctx).GetDefaultAppAccessToken(ctx, i.appCtx)
}

func NewIntegration(s *structs.AppCtx) IIntegration {
	return &Integration{appCtx: s}
}

