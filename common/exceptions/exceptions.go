package exceptions

import (
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
)

type BaseError = cExceptions.BaseError

const (
	// ErrCode_InternalError 系统错误
	ErrCode_InternalError = cExceptions.ErrCodeInternalError
	// ErrCode_InvalidParams 业务错误
	ErrCode_InvalidParams = cExceptions.ErrCodeInvalidParams
	// ErrCode_DeveloperError 开发者错误
	ErrCode_DeveloperError = cExceptions.ErrCodeDeveloperError

	// ErrCode_RATE_LIMIT_ERROR限流错误码
	ErrCode_Rate_Limit = "k_cf_ec_200009"

	// ErrCode_Flow_NoPermission 无权限操作流程
	ErrCode_Flow_NoPermission = "k_wf_ec_200036"
	// ErrCode_Flow_NotExist 流程不存在
	ErrCode_Flow_NotExist = "k_wf_ec_2001006"
	// ErrCode_Flow_ExecutionNotExist 流程实例 ID 不存在
	ErrCode_Flow_ExecutionNotExist = "k_wf_ec_2001005"
	// ErrCode_Flow_NotSupportCallFlowType 不支持调用此类型流程
	ErrCode_Flow_NotSupportCallFlowType = "k_wf_ec_2001004"
	// ErrCode_Flow_NotSupportRevokeFlow 不支持撤销该流程 (无人工任务)
	ErrCode_Flow_NotSupportRevokeFlow = "k_wf_ec_2001003"
	// ErrCode_Flow_NoReqInputParam 缺少流程的必填参数
	ErrCode_Flow_NoReqInputParam = "k_wf_ec_2001002"
	// ErrCode_Flow_InvalidParam 流程参数中 APIName 无效
	ErrCode_Flow_InvalidParam = "k_wf_ec_2001001"
)

// special for golang
const (
	// ErrCode_Data_RecordNotFound 记录不存在
	ErrCode_Data_RecordNotFound = "k_cf_ec_300004"
)

var (
	ErrTypeRecordNotFound = cExceptions.NewErrWithCode(ErrCode_Data_RecordNotFound, "data record not found")
)

func IsRecordNotFoundErr(err error) bool {
	if err == nil {
		return false
	}
	return cExceptions.ErrWrap(err).Code == ErrCode_Data_RecordNotFound
}
