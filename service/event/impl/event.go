package impl

import (
	"context"

	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type Event struct {
	appCtx *structs.AppCtx
}

func NewEvent(appCtx *structs.AppCtx) *Event {
	return &Event{appCtx: appCtx}
}

func (e *Event) GetEventInfo(ctx context.Context) (*structs.EventInfo, error) {
	return cUtils.GetEventInfoFromCtx(ctx)
}
