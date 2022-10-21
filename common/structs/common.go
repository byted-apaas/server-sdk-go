// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import "time"

type RPCCliConf struct {
	Psm         string        `yaml:"Psm" json:"Psm"`
	DebugAddr   string        `yaml:"DebugAddr" json:"DebugAddr"`
	Cluster     string        `yaml:"Cluster" json:"Cluster"`
	IDC         string        `yaml:"IDC" json:"IDC"`
	Timeout     time.Duration `yaml:"Timeout" json:"Timeout"`
	ConnTimeout time.Duration `yaml:"ConnTimeout" json:"ConnTimeout"`
}

type TimeZone struct {
	Location string `json:"location"`
	Offset   string `json:"offset"`
}

type UserSetting struct {
	UserID   int64    `json:"userID"`
	Language string   `json:"language"`
	Timezone TimeZone `json:"timezone"`
}

type Locale = UserSetting
