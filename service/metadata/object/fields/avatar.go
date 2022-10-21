// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package fields

type AvatarOrLogo struct {
	FieldBase
	DisplayStyle string `json:"displayStyle"`
}
