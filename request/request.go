// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package request

import (
	"context"
	"io"
	"sync"
	"time"

	cStructs "github.com/byted-apaas/server-common-go/structs"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request/openapi"
	"github.com/byted-apaas/server-sdk-go/service/tasks"
)

//go:generate mockery --name=IRequestOpenapi --structname=RequestOpenapi --filename=RequestOpenapi.go
type IRequestOpenapi interface {
	GetRecords(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam, records interface{}) ([][]string, error)
	GetRecordsV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2, records interface{}) ([][]string, error)
	GetRecordsV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV3, records interface{}) ([][]string, error)
	GetSingleRecordsV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, fields []string, record interface{}) ([][]string, error)
	GetRecordCount(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam) (int64, error)
	GetRecordCountV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2) (int64, error)
	GetRecordCountV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV3) (int64, error)

	CreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error)
	CreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error)
	CreateRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordIDV3, error)
	BatchCreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error)
	BatchCreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error)
	BatchCreateRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]string, error)
	BatchCreateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) (int64, error)

	UpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error
	UpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error
	UpdateRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID string, record interface{}) error
	BatchUpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) (*structs.BatchResult, error)
	BatchUpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) error
	BatchUpdateRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[string]interface{}) (*structs.BatchResultV3, error)
	BatchUpdateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) (int64, error)

	DeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error
	DeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error
	DeleteRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID string) error
	BatchDeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) (*structs.BatchResult, error)
	BatchDeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) error
	BatchDeleteRecordV3(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []string) (*structs.BatchResultV3, error)
	BatchDeleteRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) (int64, error)

	Oql(ctx context.Context, appCtx *structs.AppCtx, oql string, args interface{}, namedArgs map[string]interface{}, resultSet interface{}) (unauthFields [][]string, err error)
	Transaction(ctx context.Context, appCtx *structs.AppCtx, placeholders map[string]int64, operations []*structs.TransactionOperation, dataVersion string) (map[string]int64, error)

	DownloadFile(ctx context.Context, appCtx *structs.AppCtx, fileID string) ([]byte, error)
	DownloadAvatar(ctx context.Context, appCtx *structs.AppCtx, imageID string) ([]byte, error)
	UploadFile(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader, seconds time.Duration) (*structs.Attachment, error)
	UploadFileV2(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.FileUploadResult, error)
	UploadAvatar(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.Avatar, error)

	CreateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) (int64, error)
	UpdateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) error

	GetAllGlobalConfig(ctx context.Context, appCtx *structs.AppCtx) (map[string]string, error)

	GetFields(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string) (*structs.ObjFields, error)
	GetField(ctx context.Context, appCtx *structs.AppCtx, objectAPIName, fieldAPIName string) (*structs.Field, error)

	MGetUserSettings(ctx context.Context, appCtx *structs.AppCtx, userIDList []int64) ([]*structs.Locale, error)

	InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error
	InvokeFunctionAsync(ctx context.Context, appCtx *structs.AppCtx, apiName string, params map[string]interface{}) (int64, error)

	GetTenantInfo(ctx context.Context, appCtx *structs.AppCtx) (*cStructs.Tenant, error)

	GetExecutionUserTaskInfo(ctx context.Context, appCtx *structs.AppCtx, instanceID int64) ([]*structs.TaskInfo, error)

	Execute(ctx context.Context, appCtx *structs.AppCtx, APIName string, options *structs.ExecuteOptions) (invokeResult *structs.FlowExecuteResult, err error)
	RevokeExecution(ctx context.Context, appCtx *structs.AppCtx, instanceID int64, options *structs.RevokeOptions) error
	GetExecutionInfo(ctx context.Context, appCtx *structs.AppCtx, instanceID int64) (*structs.ExecutionInfo, error)
	GetTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.TenantAccessToken, error)
	GetAppAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.AppAccessToken, error)
	GetDefaultTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.TenantAccessToken, error)
	GetDefaultAppAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.AppAccessToken, error)
	GetApprovalInstanceList(ctx context.Context, appCtx *structs.AppCtx, options *structs.ApprovalInstanceListOptions) (*structs.ApprovalInstanceList, error)
	GetApprovalInstance(ctx context.Context, appCtx *structs.AppCtx, options *structs.GetApprovalInstanceOptions) (*structs.ApprovalInstance, error)
}

//go:generate mockery --name=IRequestOpenapi --structname=RequestOpenapi --filename=RequestOpenapi.go
type IRequestFaaSInfra interface {
	InvokeFunction(ctx context.Context, appCtx *structs.AppCtx, apiName string, params map[string]interface{}, result interface{}) error
	InvokeFunctionAsync(ctx context.Context, appCtx *structs.AppCtx, apiName string, params map[string]interface{}) (int64, error)
	InvokeFunctionDistributed(ctx context.Context, appCtx *structs.AppCtx, dataset interface{}, handlerFunc string, progressCallbackFunc string, completedCallbackFunc string, options *tasks.Options) (int64, error)
}

var (
	reqHttp     IRequestOpenapi
	reqHttpOnce sync.Once
)

func GetInstance(ctx context.Context) IRequestOpenapi {
	return GetHttpInstance()
}

func GetHttpInstance() IRequestOpenapi {
	if reqHttp == nil {
		reqHttpOnce.Do(func() {
			reqHttp = &openapi.RequestHttp{}
		})
	}
	return reqHttp
}
