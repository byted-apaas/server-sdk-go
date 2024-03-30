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
	Content interface{}  `json:"content"`
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
