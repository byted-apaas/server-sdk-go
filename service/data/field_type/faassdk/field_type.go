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

// AvatarV3 头像（datav3 协议）
// @example
// {
//    "source": "image",
//    "image": {
//        "token": "0a4ca526dxxxxxxxxx9457145882",                    // 附件ID， 富文本resource id 或 图片ID
//        "uri": "/img/198999/da69a4ef2ebxxxxxxxxba397bcc6602_l.jpg", // 图片large信息
//    }
//}
type AvatarV3 struct {
	Source  string             `json:"source"`
	Image   *AttachmentModelV3 `json:"image,omitempty"`
	Color   *string            `json:"color,omitempty"`
	Content interface{}        `json:"content,omitempty"` // int
}

// AttachmentModelV3 文件（datav3 协议）
type AttachmentModelV3 struct {
	Token    string `json:"token"`
	MimeType string `json:"mime_type,omitempty"`
	Name     string `json:"name,omitempty"`
	Size     string `json:"size,omitempty"`
	URI      string `json:"uri,omitempty"`
}

// RichTextV3 富文本（datav3 协议）
type RichTextV3 struct {
	Raw    string             `json:"raw"`
	Config *AttachmentModelV3 `json:"config,omitempty"`
}

// PhoneNumberV3 DataV3 协议
type PhoneNumberV3 struct {
	RegionCode  string `json:"region_code"`
	DialingCode string `json:"dialing_code"`
	Number      string `json:"number"`
}

type RegionV3 struct {
	ID         string         `json:"_id"`
	RegionCode string         `json:"region_code"`
	FullPath   MultilingualV3 `json:"full_path"`
}

type MultilingualV3 struct {
	En string `json:"en_US,omitempty"`
	Zh string `json:"zh_CN,omitempty"`
}
