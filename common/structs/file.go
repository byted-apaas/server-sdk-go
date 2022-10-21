// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

type FileUploadResult struct {
	FileID string `json:"fileId"`
	Type   string `json:"type"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
}

type Attachment struct {
	ID       string `json:"id"`
	MimeType string `json:"type"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Token    string `json:"token"`
}

type Avatar struct {
	ImageID        string `json:"imageId"`
	PreviewImageID string `json:"previewImageId"`
}
