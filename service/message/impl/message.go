// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package impl

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
	"github.com/byted-apaas/server-sdk-go/service/message"
)

type Msg struct {
	NotifyCenter message.INotifyCenter
}

func NewMsg(s *structs.AppCtx) *Msg {
	return &Msg{
		NotifyCenter: NewNotifyCenter(s),
	}
}

type NotifyCenter struct {
	appCtx *structs.AppCtx
}

func NewNotifyCenter(s *structs.AppCtx) message.INotifyCenter {
	return &NotifyCenter{appCtx: s}
}

func (n *NotifyCenter) Create(ctx context.Context, msg *structs.MessageBody) (int64, error) {
	param, err := msgBodyToParam(ctx, msg)
	if err != nil {
		return 0, err
	}

	return request.GetInstance(ctx).CreateMessage(ctx, n.appCtx, param)
}

func (n *NotifyCenter) Update(ctx context.Context, msgID int64, msg *structs.MessageBody) error {
	param, err := msgBodyToParam(ctx, msg)
	if err != nil {
		return err
	}
	param["TaskID"] = msgID

	return request.GetInstance(ctx).UpdateMessage(ctx, n.appCtx, param)
}

func msgBodyToParam(ctx context.Context, msg *structs.MessageBody) (map[string]interface{}, error) {
	if msg == nil {
		return nil, cExceptions.InvalidParamError("The msg can not be empty")
	}

	body := map[string]interface{}{}

	rawData := map[string]interface{}{}
	searchData := map[string]interface{}{}
	if !cUtils.StrInStrs(constants.MsgIconTypes, msg.Icon) {
		return nil, cExceptions.InvalidParamError("The value of msg.icon is invalid")
	} else {
		s, err := cUtils.JsonMarshalBytes(msg.Icon)
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.icon failed, err: %+v", err)
		}
		rawData["icon"] = string(s)
		searchData["icon_type"] = msg.Icon
	}

	if msg.Percent < 0 || msg.Percent > 100 {
		return nil, cExceptions.InvalidParamError("Invalid value of parameter percent, the value should be in range [0, 100]")
	} else {
		s, err := cUtils.JsonMarshalBytes(msg.Percent)
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.percent failed, err: %+v", err)
		}
		rawData["percent"] = string(s)
		searchData["percent"] = string(s)
	}

	if len(msg.TargetUsers) == 0 {
		return nil, cExceptions.InvalidParamError("Marshal msg.targetUsers can not be empty")
	} else {
		s, err := cUtils.JsonMarshalBytes(msg.TargetUsers)
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.targetUsers failed, err: %+v", err)
		}
		rawData["target_users"] = string(s)
	}

	if msg.Title != nil {
		t, err := cUtils.JsonMarshalBytes(msg.Title)
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.title failed, err: %+v", err)
		}

		s, err := cUtils.JsonMarshalBytes(map[string]interface{}{"value": string(t)})
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.title failed, err: %+v", err)
		}
		rawData["title"] = string(s)
	}

	if len(msg.TargetUsers) == 0 {
		return nil, cExceptions.InvalidParamError("Marshal msg.targetUsers can not be empty")
	} else {
		d, err := cUtils.JsonMarshalBytes(msg.Detail)
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.detail failed, err: %+v", err)
		}
		s, err := cUtils.JsonMarshalBytes(map[string]interface{}{"value": string(d)})
		if err != nil {
			return nil, cExceptions.InvalidParamError("Marshal msg.detail failed, err: %+v", err)
		}
		rawData["detail"] = string(s)
	}

	body["NotifyModelKey"] = constants.MsgSourcePSM
	body["ReceiverIDs"] = msg.TargetUsers
	body["Percent"] = msg.Percent
	body["ParamsRawData"] = rawData
	body["ModelKeySearchData"] = searchData
	return body, nil
}
