package event

import (
	"context"
)

type IEvent interface {
	GetEventInfo(ctx context.Context) (*cStructs.EventInfo, error)
}
