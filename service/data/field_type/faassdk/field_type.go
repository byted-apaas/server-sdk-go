// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package faassdk

import "github.com/byted-apaas/server-sdk-go/common/structs/intern"

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
