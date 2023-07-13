package event

import (
	"context"

	cStructs "github.com/byted-apaas/server-common-go/structs"
)

type IEvent interface {
	GetEventInfo(ctx context.Context) (*cStructs.EventInfo, error)
}
