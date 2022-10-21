// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package user

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

type IUser interface {
	// GetUserID 若无实际触发人, 返回 -1
	GetUserID(ctx context.Context) int64
	// GetLocaleByUserID 获取指定 user 的 locale 信息
	GetLocaleByUserID(ctx context.Context, userID int64) (*structs.Locale, error)
	// GetLocaleByUserIDList 批量获取指定 user 列表的 locale 信息
	GetLocaleByUserIDList(ctx context.Context, userIDList []int64) ([]*structs.Locale, error)
}

type User struct {
	appCtx *structs.AppCtx
}

func NewUser(s *structs.AppCtx) IUser {
	return &User{appCtx: s}
}

func (u *User) GetUserID(ctx context.Context) int64 {
	return cUtils.GetUserIDFromCtx(ctx)
}

func (u *User) GetLocaleByUserID(ctx context.Context, userID int64) (*structs.Locale, error) {
	localeList, err := u.GetLocaleByUserIDList(ctx, []int64{userID})
	if err != nil {
		return nil, err
	}
	if len(localeList) < 1 {
		return nil, cExceptions.InvalidParamError("locale is empty for userID(%v)", userID)
	}
	return localeList[0], nil
}

func (u *User) GetLocaleByUserIDList(ctx context.Context, userIDList []int64) ([]*structs.Locale, error) {
	settings, err := request.GetInstance(ctx).MGetUserSettings(ctx, u.appCtx, userIDList)
	if err != nil {
		return nil, err
	}
	return settings, nil
}
