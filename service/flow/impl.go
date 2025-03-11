package flow

import (
	"context"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

type Flow struct {
	appCtx *structs.AppCtx
}

func NewFlow(appCtx *structs.AppCtx) IFlow {
	return &Flow{appCtx: appCtx}
}

func (w *Flow) Execute(ctx context.Context, APIName string, options *structs.ExecuteOptions, async *bool) (invokeResult *structs.FlowExecuteResult, err error) {
	if APIName == "" {
		return nil, cExceptions.InvalidParamError("APIName is empty")
	}
	if options == nil {
		options = &structs.ExecuteOptions{}
	}
	return request.GetInstance(ctx).Execute(ctx, w.appCtx, APIName, options, async)
}

func (w *Flow) RevokeExecution(ctx context.Context, executionID int64, options *structs.RevokeOptions) error {
	return request.GetInstance(ctx).RevokeExecution(ctx, w.appCtx, executionID, options)
}

func (w *Flow) GetExecutionInfo(ctx context.Context, executionID int64) (*structs.ExecutionInfo, error) {
	return request.GetInstance(ctx).GetExecutionInfo(ctx, w.appCtx, executionID)
}

func (w *Flow) GetExecutionUserTaskInfo(ctx context.Context, executionID int64) ([]*structs.TaskInfo, error) {
	return request.GetInstance(ctx).GetExecutionUserTaskInfo(ctx, w.appCtx, executionID)
}

func (w *Flow) GetApprovalInstanceList(ctx context.Context, options *structs.ApprovalInstanceListOptions) (*structs.ApprovalInstanceList, error) {
	return request.GetInstance(ctx).GetApprovalInstanceList(ctx, w.appCtx, options)
}

func (w *Flow) GetApprovalInstance(ctx context.Context, options *structs.GetApprovalInstanceOptions) (*structs.ApprovalInstance, error) {
	return request.GetInstance(ctx).GetApprovalInstance(ctx, w.appCtx, options)
}
