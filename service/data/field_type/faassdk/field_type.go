// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package faassdk

const (
	ZH = 2052
	EN = 1033
)

// Multilingual 多语结构
type Multilingual []*MultilingualItem

type MultilingualItem struct {
	LanguageCode int    `json:"language_code"`
	Text         string `json:"text"`
}

// PhoneNumber 电话号码
type PhoneNumber struct {
	// 示例：+86(CN)
	Key string `json:"key"`
	// 示例：18898888888
	Number string `json:"number"`
}
