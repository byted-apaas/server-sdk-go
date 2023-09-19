package exceptions

import (
	"fmt"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
)

type BaseError = cExceptions.BaseError

const (
	// ErrCode_InternalError 系统错误
	ErrCode_InternalError = cExceptions.ErrCodeInternalError
	// ErrCode_InvalidParams 业务错误
	ErrCode_InvalidParams = cExceptions.ErrCodeDeveloperError
	// ErrCode_DeveloperError 开发者错误
	ErrCode_DeveloperError = cExceptions.ErrCodeDeveloperError

	// ErrCode_RATE_LIMIT_ERROR限流错误码
	ErrCode_Rate_Limit = cExceptions.ErrCodeRateLimitError

	// ErrCode_Flow_NoPermission 无权限操作流程
	ErrCode_Flow_NoPermission = cExceptions.ErrCodeFlowNoPermission
	// ErrCode_Flow_NotExist 流程不存在
	ErrCode_Flow_NotExist = cExceptions.ErrCodeFlowNotExist
	// ErrCode_Flow_ExecutionNotExist 流程实例 ID 不存在
	ErrCode_Flow_ExecutionNotExist = cExceptions.ErrCodeFlowExecutionNotExist
	// ErrCode_Flow_NotSupportCallFlowType 不支持调用此类型流程
	ErrCode_Flow_NotSupportCallFlowType = cExceptions.ErrCodeFlowNotSupportCallFlowType
	// ErrCode_Flow_NotSupportRevokeFlow 不支持撤销该流程 (无人工任务)
	ErrCode_Flow_NotSupportRevokeFlow = cExceptions.ErrCodeFlowNotSupportRevokeFlow
	// ErrCode_Flow_NoReqInputParam 缺少流程的必填参数
	ErrCode_Flow_NoReqInputParam = cExceptions.ErrCodeFlowNoReqInputParam
	// ErrCode_Flow_InvalidParam 流程参数中 APIName 无效
	ErrCode_Flow_InvalidParam = cExceptions.ErrCodeFlowInvalidParam
)

// special for golang
const (
	// ErrCode_Data_RecordNotFound 记录不存在
	ErrCode_Data_RecordNotFound = cExceptions.ErrCodeDataRecordNotFound
)

var (
	ErrTypeRecordNotFound = cExceptions.NewErrWithCode(cExceptions.ErrCodeDataRecordNotFound, "data record not found", "")
)

func IsRecordNotFoundErr(err error) bool {
	if err == nil {
		return false
	}
	return cExceptions.ErrWrap(err).Code == ErrCode_Data_RecordNotFound
}

var (
	ErrRecordIsEmpty     = fmt.Errorf("record is empty")
	ErrFieldNotFound     = fmt.Errorf("field not found")
	ErrFieldTypeNotMatch = fmt.Errorf("field type not match")
	ErrFieldIsEmpty      = fmt.Errorf("field is empty")
	ErrFieldNoPermission = fmt.Errorf("field no permission")
)
