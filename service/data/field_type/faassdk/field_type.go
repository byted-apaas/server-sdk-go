// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package faassdk

import (
	"github.com/byted-apaas/server-sdk-go/common/structs/intern"
)

const (
	ZH = intern.LanguageCodeZH
	EN = intern.LanguageCodeEN
)

// Multilingual 多语结构
type Multilingual = intern.Multilingual

// PhoneNumber 电话号码
type PhoneNumber struct {
	// 示例：+86(CN)
	Key string `json:"key"`
	// 示例：18898888888
	Number string `json:"number"`
}

// Region 行政区划
type Region struct {
	ID         int64        `json:"id"`
	RegionID   int64        `json:"regionId"`
	Level      int          `json:"level"`
	RegionCode string       `json:"regionCode"`
	FullPath   Multilingual `json:"fullPath"`
}

type AvatarImages map[string]string

// Avatar 头像
type Avatar struct {
	Source  string       `json:"source"`
	Image   AvatarImages `json:"image"`
	Color   string       `json:"color"`
	ColorID string       `json:"color_id"`
	Content int64        `json:"content"` // todo wby 确认这个字段类型：*interface{}
}

// RichTextConfig 富文本配置
type RichTextConfig struct {
	ResourceID   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Token        string `json:"token"`
}

// RichText 富文本
type RichText struct {
	Raw     string            `json:"raw"`
	Preview string            `json:"preview"`
	Config  []*RichTextConfig `json:"config"`
}

// Attachment 文件
type Attachment struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	MimeType string `json:"mime_type"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
}

// AvatarV2 头像 todo wby
// @example
// {
//    "source": "image",
//    "image": {
//        "token": "0a4ca526dxxxxxxxxx9457145882",                    // 附件ID， 富文本resource id 或 图片ID
//        "uri": "/img/198999/da69a4ef2ebxxxxxxxxba397bcc6602_l.jpg", // 图片large信息
//    }
//}
type AvatarV2 struct {
	Source string       `json:"source"`
	Image  AvatarImages `json:"image"`
}
