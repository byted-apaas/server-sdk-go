// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package message

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

//go:generate mockery --name=INotifyCenter --structname=NotifyCenter --filename=NotifyCenter.go
type INotifyCenter interface {
	Create(ctx context.Context, msg *structs.MessageBody) (msgID int64, err error)
	Update(ctx context.Context, msgID int64, msg *structs.MessageBody) (err error)
}
