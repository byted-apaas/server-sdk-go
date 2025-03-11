package flow

import (
	"context"

	"github.com/byted-apaas/server-sdk-go/common/structs"
)

type IFlow interface {
	// GetExecutionUserTaskInfo 获取流程实例人工任务详情列表
	GetExecutionUserTaskInfo(ctx context.Context, executionID int64) (taskInfo []*structs.TaskInfo, err error)

	// Execute 触发流程 (需要关注目前能够触发的流程类型，见使用指南)
	Execute(ctx context.Context, APIName string, options *structs.ExecuteOptions, async *bool) (invokeResult *structs.FlowExecuteResult, err error)
	// RevokeExecution 撤销流程 (仅支持撤销包含人工任务的流程)
	RevokeExecution(ctx context.Context, executionID int64, options *structs.RevokeOptions) (err error)
	// GetExecutionInfo 获取流程实例信息
	GetExecutionInfo(ctx context.Context, executionID int64) (info *structs.ExecutionInfo, err error)
	GetApprovalInstanceList(ctx context.Context, options *structs.ApprovalInstanceListOptions) (*structs.ApprovalInstanceList, error)
	GetApprovalInstance(ctx context.Context, options *structs.GetApprovalInstanceOptions) (*structs.ApprovalInstance, error)
}
