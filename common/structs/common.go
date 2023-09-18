// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import (
	"context"
	"time"
)

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
	UserID   int64    `json:"userId"`
	Language string   `json:"language"`
	Timezone TimeZone `json:"timezone"`
}

type Locale = UserSetting

type Multilingual struct {
	En string `json:"en_US,omitempty"`
	Zh string `json:"zh_CN,omitempty"`
}

type FlowUser struct {
	// aPaaS UserID
	UserID int64 `json:"userId"`
	// Lark OpenID
	FeishuOpenID string `json:"feishuOpenId"`
	// 用户名
	UserName Multilingual `json:"userName"`
	// 用户头像
	//UserAvatar FlowUserAvatar `json:"userAvatar"`
}

type FlowUserAvatar struct {
	// 用户头像 URL
	URL string `json:"url"`
}

type TaskInfo struct {
	TaskID int64 `json:"taskId"`
	// 任务名称
	Label Multilingual `json:"label"`
	// 任务类型
	TaskType string `json:"taskType"`
	// 任务状态
	TaskStatus TaskStatus `json:"taskStatus"`
	// 任务意见列表
	OpinionList []*TaskOpinion `json:"opinionList"`
	// 指派人列表
	AssigneeList []*FlowUser `json:"assigneeList"`
	// 节点开始时间
	StartedTime int64 `json:"startedTime"`
	// 节点完成时间
	CompletedTime *int64 `json:"completedTime"`
}

type TaskOpinion struct {
	User FlowUser `json:"user"`
	// 意见结果
	OpinionResult string `json:"opinionResult"`
	// 提交时间
	SubmitTime int64 `json:"submitTime"`
	// 意见内容
	OpinionContent *string `json:"opinionContent"`
}

type TaskType string

const (
	Form     TaskType = "form"
	Approval TaskType = "approval"
	CC       TaskType = "CC"
)

type TaskStatus string

const (
	// TaskStatusInProcess 等待审批中
	TaskStatusInProcess = "in_process"
	// TaskStatusCompleted 任务完成
	TaskStatusCompleted = "completed"
	// TaskStatusCanceled 任务取消
	TaskStatusCanceled = "canceled"
	// TaskStatusFailed 任务失败
	TaskStatusFailed = "failed"
	// TaskStatusAutoEnd 自动结束
	TaskStatusAutoEnd = "auto_end"
	// TaskStatusAgreed 用户通过
	TaskStatusAgreed = "agreed"
	// TaskStatusRejected 用户拒绝
	TaskStatusRejected = "rejected"
)

type ExecuteOptions struct {
	// 流程入参
	Params map[string]interface{} `json:"params"`
}

type FlowExecuteResult struct {
	ExecutionID int64                  `json:"executionId"`
	Status      ExecutionStatus        `json:"status"`
	Data        map[string]interface{} `json:"data"`
	ErrCode     *string                `json:"errCode"`
	ErrMsg      *string                `json:"errMsg"`
}

type RevokeOptions struct {
	Reason Multilingual `json:"reason"`
}
type ExecutionInfo struct {
	Status  ExecutionStatus        `json:"status"`
	Data    map[string]interface{} `json:"data"`
	ErrCode *string                `json:"errCode"`
	ErrMsg  *string                `json:"errMsg"`
}
type ExecutionStatus string

const (
	ExecutionStatusWait      ExecutionStatus = "wait"
	ExecutionStatusInProcess ExecutionStatus = "in_process"
	ExecutionStatusEnd       ExecutionStatus = "end"
	ExecutionStatusRejected  ExecutionStatus = "rejected"
	ExecutionStatusCanceled  ExecutionStatus = "canceled"
	ExecutionStatusFailed    ExecutionStatus = "failed"
	ExecutionStatusAgreed    ExecutionStatus = "agreed"
)

type TenantAccessToken struct {
	Expire            int64  `json:"expire"`
	TenantAccessToken string `json:"tenantAccessToken"`
	AppID             string `json:"appId"`
}

type AppAccessToken struct {
	Expire         int64  `json:"expire"`
	AppAccessToken string `json:"appAccessToken"`
	AppID          string `json:"appId"`
}

type FindStreamData struct {
	Records      interface{} `json:"records"`
	UnauthFields [][]string  `json:"unauthFields"`
}

type FindStreamParam struct {
	IDGetter func(record interface{}) (id int64, err error)
	Handler  func(ctx context.Context, data *FindStreamData) (err error)
}
