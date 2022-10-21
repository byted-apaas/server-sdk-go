// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package opensdk

type Multilingual struct {
	En string `json:"en_US,omitempty"`
	Zh string `json:"zh_CN,omitempty"`
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	DialingCode string `json:"dialingCode"`
	Number      string `json:"number"`
}

type File struct {
	FileID string `json:"fileId"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Type   string `json:"type"`
}

type RichTextConfig struct {
	ResourceID   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
}

type RichText struct {
	Config []*RichTextConfig `json:"config"`
	Raw    string            `json:"raw"`
}
