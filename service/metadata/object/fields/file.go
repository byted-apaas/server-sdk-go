// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type File struct {
	FieldBase
	Required  bool     `json:"required"`
	Multiple  bool     `json:"multiple"`
	FileTypes []string `json:"fileTypes"`
}

type FileV3 struct {
	FieldBaseV3
	Required  bool     `json:"required"`
	Multiple  bool     `json:"multiple"`
	FileTypes []string `json:"fileTypes"`
}
