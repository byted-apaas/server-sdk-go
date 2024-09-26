// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package openapi

import (
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

const (
	// DataV3 接口
	PathCreateRecordV3      = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records"
	PathBatchCreateRecordV3 = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records_batch"

	PathDeleteRecordV3      = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records/:recordID"
	PathBatchDeleteRecordV3 = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records_batch"

	PathUpdateRecordV3      = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records/:recordID"
	PathBatchUpdateRecordV3 = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records_batch"

	PathGetRecordsV3      = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records/:recordID"
	PathBatchGetRecordsV3 = "/v1/data/namespaces/:namespace/objects/:objectAPIName/records_query"

	// PathGetRecordsV2 新版接口
	PathGetRecordsV2 = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/records"

	PathCreateRecordV2      = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName"
	PathBatchCreateRecordV2 = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/batchCreate"

	PathUpdateRecordV2      = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/:recordID"
	PathBatchUpdateRecordV2 = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/batchUpdate"

	PathDeleteRecordV2      = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/:recordID"
	PathBatchDeleteRecordV2 = "/api/data/v1/namespaces/:namespace/objects/:objectAPIName/batchDelete"

	PathUploadFileV2   = "/api/attachment/v1/files"
	PathUploadAvatar   = "/api/attachment/v1/images"
	PathDownloadFileV2 = "/api/attachment/v1/files/:fileID"
	PathDownloadAvatar = "/api/attachment/v1/images/:fileID"

	// PathGetRecords 老版接口
	PathGetRecords  = "/generic/data/v3/namespaces/:namespace/objects/:objectAPIName/mGetByCriterion"
	PathTransaction = "/generic/datax/v2/namespaces/:namespace/records/_modify_with_transaction"
	PathOql         = "/data/v5/namespaces/:namespace/records/query"

	PathCreateRecord           = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/createBySync"
	PathBatchCreateRecord      = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchCreate"
	PathBatchCreateRecordAsync = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchCreateByAsync"

	PathUpdateRecord           = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/updateBySync"
	PathBatchUpdateRecord      = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchUpdate"
	PathBatchUpdateRecordAsync = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchUpdateByAsync"

	PathDeleteRecord           = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/deleteBySync"
	PathBatchDeleteRecord      = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchDelete"
	PathBatchDeleteRecordAsync = "/generic/data/v2/namespaces/:namespace/objects/:objectAPIName/batchDeleteByAsync"

	PathUploadFile   = "/attachment/v2/upload"
	PathDownloadFile = "/attachment/v1/download"

	PathCreateMessage = "/msg/v4/namespaces/:namespace/message_center/notify_task/create"
	PathUpdateMessage = "/msg/v4/namespaces/:namespace/message_center/notify_task/update"

	PathGetAllGlobalConfig = "/generic/globalconfig/v1/namespaces/:namespace/getAllConfigs"

	PathGetField  = "/metadata/v5/namespaces/:namespace/objects/:objectAPIName/fields/:fieldAPIName"
	PathGetFields = "/metadata/v5/namespaces/:namespace/objects/:objectAPIName/describe"

	PathMGetUserSettings = "/setting/v1/users"

	PathInvokeFunctionWithAuth = "/api/cloudfunction/v1/namespaces/:namespace/invoke/:functionAPIName"
	PathInvokeFunctionAsync    = "/faasAsyncTask/v1/namespaces/:namespace/asyncTask/CreateAsyncTask"

	PathGetTaskListByExecutionID = "/api/flow/v1/namespaces/:namespace/executions/:executionId/userTaskInfo"
	PathGetExecutionInfo         = "/api/flow/v1/namespaces/:namespace/executions/:executionId/detail"
	PathRevokeExecution          = "/api/flow/v1/namespaces/:namespace/executions/:executionId/revoke"
	PathExecuteFlow              = "/api/flow/v1/namespaces/:namespace/flows/:apiName/execute"
	PathDefaultTenantAccessToken = "/api/integration/v1/namespaces/:namespace/defaultLark/tenantAccessToken"
	PathDefaultAppAccessToken    = "/api/integration/v1/namespaces/:namespace/defaultLark/appAccessToken"
	PathAppAccessToken           = "/api/integration/v1/namespaces/:namespace/lark/appAccessToken/:apiName"
	PathTenantAccessToken        = "/api/integration/v1/namespaces/:namespace/lark/tenantAccessToken/:apiName"
	PathGetApprovalInstanceList  = "/api/approval/v1/approval_instances/listids"
	PathGetApprovalInstance      = "/api/approval/v1/approval_instances/:recordID"
)

// GetPathGetRecordsV2 新版接口
func GetPathGetRecordsV2(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathGetRecordsV2).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathGetRecordsV3(namespace, objectAPIName string, recordID int64) string {
	return cUtils.NewPathReplace(PathGetRecordsV3).Namespace(namespace).ObjectAPIName(objectAPIName).RecordID(recordID).Path()
}

func GetPathBatchGetRecordsV3(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchGetRecordsV3).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathCreateRecordV2(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathCreateRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathCreateRecordV3(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathCreateRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchCreateRecordV3(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchCreateRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchCreateRecordV2(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchCreateRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathUpdateRecordV2(namespace, objectAPIName string, recordID int64) string {
	return cUtils.NewPathReplace(PathUpdateRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).RecordID(recordID).Path()
}

func GetPathUpdateRecordV3(namespace, objectAPIName string, recordID string) string {
	return cUtils.NewPathReplace(PathUpdateRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).RecordIDStr(recordID).Path()
}

func GetPathBatchUpdateRecordV3(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchUpdateRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchUpdateRecordV2(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchUpdateRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathDeleteRecordV2(namespace, objectAPIName string, recordID int64) string {
	return cUtils.NewPathReplace(PathDeleteRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).RecordID(recordID).Path()
}

func GetPathBatchDeleteRecordV2(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchDeleteRecordV2).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathDeleteRecordV3(namespace, objectAPIName string, recordID string) string {
	return cUtils.NewPathReplace(PathDeleteRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).RecordIDStr(recordID).Path()
}

func GetPathBatchDeleteRecordV3(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchDeleteRecordV3).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathDownloadFileV2(fileID string) string {
	return cUtils.NewPathReplace(PathDownloadFileV2).FileID(fileID).Path()
}

func GetPathDownloadAvatar(imageID string) string {
	return cUtils.NewPathReplace(PathDownloadAvatar).FileID(imageID).Path()
}

// GetPathGetRecords 老版接口
func GetPathGetRecords(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathGetRecords).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathTransaction(namespace string) string {
	return cUtils.NewPathReplace(PathTransaction).Namespace(namespace).Path()
}

func GetPathOql(namespace string) string {
	return cUtils.NewPathReplace(PathOql).Namespace(namespace).Path()
}

func GetPathCreateRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathCreateRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchCreateRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchCreateRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchCreateRecordAsync(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchCreateRecordAsync).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathUpdateRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathUpdateRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchUpdateRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchUpdateRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchUpdateRecordAsync(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchUpdateRecordAsync).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathDeleteRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathDeleteRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchDeleteRecord(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchDeleteRecord).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathBatchDeleteRecordAsync(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathBatchDeleteRecordAsync).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathDownloadFile() string {
	return cUtils.NewPathReplace(PathDownloadFile).Path()
}

func GetPathCreateMessage(namespace string) string {
	return cUtils.NewPathReplace(PathCreateMessage).Namespace(namespace).Path()
}

func GetPathUpdateMessage(namespace string) string {
	return cUtils.NewPathReplace(PathUpdateMessage).Namespace(namespace).Path()
}

func GetPathGetAllGlobalConfig(namespace string) string {
	return cUtils.NewPathReplace(PathGetAllGlobalConfig).Namespace(namespace).Path()
}

func GetPathGetFields(namespace, objectAPIName string) string {
	return cUtils.NewPathReplace(PathGetFields).Namespace(namespace).ObjectAPIName(objectAPIName).Path()
}

func GetPathGetField(namespace, objectAPIName, fieldAPIName string) string {
	return cUtils.NewPathReplace(PathGetField).Namespace(namespace).ObjectAPIName(objectAPIName).FieldAPIName(fieldAPIName).Path()
}

func GetPathMGetUserSettings() string {
	return cUtils.NewPathReplace(PathMGetUserSettings).Path()
}

func GetPathInvokeFunctionWithAuth(namespace, functionAPIName string) string {
	return cUtils.NewPathReplace(PathInvokeFunctionWithAuth).Namespace(namespace).FunctionAPIName(functionAPIName).Path()
}

func GetPathInvokeFunctionAsync(namespace string) string {
	return cUtils.NewPathReplace(PathInvokeFunctionAsync).Namespace(namespace).Path()
}

func GetExecutionUserTaskInfo(namespace string, executionID int64) string {
	return cUtils.NewPathReplace(PathGetTaskListByExecutionID).Namespace(namespace).ExecutionID(executionID).Path()
}

func GetPathGetExecutionInfo(namespace string, executionID int64) string {
	return cUtils.NewPathReplace(PathGetExecutionInfo).Namespace(namespace).ExecutionID(executionID).Path()
}

func GetPathRevokeExecution(namespace string, executionID int64) string {
	return cUtils.NewPathReplace(PathRevokeExecution).Namespace(namespace).ExecutionID(executionID).Path()
}

func GetPathExecuteFlow(namespace, APIName string) string {
	return cUtils.NewPathReplace(PathExecuteFlow).Namespace(namespace).APIName(APIName).Path()
}

func GetDefaultAppAccessTokenPath(namespace string) string {
	return cUtils.NewPathReplace(PathDefaultAppAccessToken).Namespace(namespace).Path()
}

func GetDefaultTenantAccessTokenPath(namespace string) string {
	return cUtils.NewPathReplace(PathDefaultTenantAccessToken).Namespace(namespace).Path()
}

func GetAppAccessTokenPath(namespace, apiName string) string {
	return cUtils.NewPathReplace(PathAppAccessToken).Namespace(namespace).APIName(apiName).Path()
}

func GetTenantAccessTokenPath(namespace, apiName string) string {
	return cUtils.NewPathReplace(PathTenantAccessToken).Namespace(namespace).APIName(apiName).Path()
}

func GetGetApprovalInstancePath(instanceID int64) string {
	return cUtils.NewPathReplace(PathGetApprovalInstance).RecordID(instanceID).Path()
}
