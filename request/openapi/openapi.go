// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cHttp "github.com/byted-apaas/server-common-go/http"
	cStructs "github.com/byted-apaas/server-common-go/structs"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/common/structs/intern"
	"github.com/byted-apaas/server-sdk-go/common/utils"
	reqCommon "github.com/byted-apaas/server-sdk-go/request/common"
	"github.com/tidwall/gjson"
)

type RequestHttp struct{}

func (r *RequestHttp) Execute(ctx context.Context, appCtx *structs.AppCtx, APIName string, options *structs.ExecuteOptions) (invokeResult *structs.FlowExecuteResult, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.ExecuteFlow)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	body := map[string]interface{}{
		"variables": utils.ParseMapToFlowVariable(options.Params),
		"operator":  cUtils.GetUserIDFromCtx(ctx),
		"loopMasks": cUtils.GetLoopMaskFromCtx(ctx),
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathExecuteFlow(namespace, APIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var inRes intern.FlowExecuteResult
	err = cUtils.JsonUnmarshalBytes(data, &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[execute] failed, err: %v", err)
	}
	return &structs.FlowExecuteResult{
		ExecutionID: inRes.ExecutionID,
		Status:      structs.ExecutionStatus(inRes.Status),
		Data:        utils.ParseFlowVariableToMap(inRes.OutParams),
		ErrCode:     inRes.ErrCode,
		ErrMsg:      inRes.ErrMsg,
	}, nil
}

func (r *RequestHttp) RevokeExecution(ctx context.Context, appCtx *structs.AppCtx, executionID int64, options *structs.RevokeOptions) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.RevokeExecution)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"operator": cUtils.GetUserIDFromCtx(ctx),
		"reason":   options.Reason,
	}

	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathRevokeExecution(namespace, executionID), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestHttp) GetExecutionInfo(ctx context.Context, appCtx *structs.AppCtx, executionID int64) (instanceInfo *structs.ExecutionInfo, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetExecutionInfo)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetPathGetExecutionInfo(namespace, executionID)+fmt.Sprintf("?operator=%v", cUtils.GetUserIDFromCtx(ctx)), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	raw := gjson.GetBytes(data, "executionInfo").Raw

	var inRes intern.ExecutionInfo
	err = cUtils.JsonUnmarshalBytes([]byte(raw), &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[GetExecutionInfo] failed, err: %v", err)
	}

	return &structs.ExecutionInfo{
		Status:  structs.ExecutionStatus(inRes.Status),
		Data:    utils.ParseFlowVariableToMap(inRes.OutParams),
		ErrCode: inRes.ErrCode,
		ErrMsg:  inRes.ErrMsg,
	}, nil
}

func (r *RequestHttp) GetExecutionUserTaskInfo(ctx context.Context, appCtx *structs.AppCtx, executionID int64) (taskInfoList []*structs.TaskInfo, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetExecutionUserTaskInfo)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetExecutionUserTaskInfo(namespace, executionID)+fmt.Sprintf("?operator=%v", cUtils.GetUserIDFromCtx(ctx)), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	raw := gjson.GetBytes(data, "taskList").Raw
	err = cUtils.JsonUnmarshalBytes([]byte(raw), &taskInfoList)
	if err != nil {
		return nil, cExceptions.InternalError("[GetExecutionUserTaskInfo] result decode failed: %v", err)
	}
	return
}

func (r *RequestHttp) GetTenantInfo(ctx context.Context, appCtx *structs.AppCtx) (*cStructs.Tenant, error) {
	return appCtx.Credential.GetTenantInfo(utils.SetCtx(ctx, appCtx, cConstants.GetAppToken))
}

func (r *RequestHttp) GetRecords(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam, records interface{}) error {
	if param == nil {
		param = &structs.GetRecordsReqParam{}
	}
	if param.Limit <= 0 {
		param.Limit = constants.PageLimitMax
	}
	param.NeedFilterUserPermission = false
	param.IgnoreBackLookupField = false

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathGetRecords(namespace, objectAPIName), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return err
	}

	recordsRaw := gjson.GetBytes(data, "data_list").Raw
	err = cUtils.JsonUnmarshalBytes([]byte(recordsRaw), records)
	if err != nil {
		return cExceptions.InvalidParamError("GetRecords failed, err: %v", err)
	}

	return nil
}

func (r *RequestHttp) GetRecordsV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2, records interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetRecordsV2)

	if param == nil {
		param = &structs.GetRecordsReqParamV2{}
	}
	if param.Limit <= 0 {
		param.Limit = constants.PageLimitMax
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathGetRecordsV2(namespace, objectAPIName), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return err
	}

	recordsRaw := gjson.GetBytes(data, "records").Raw
	err = cUtils.JsonUnmarshalBytes([]byte(recordsRaw), records)
	if err != nil {
		return cExceptions.InvalidParamError("GetRecordsV2 failed, err: %v", err)
	}

	return nil
}

func (r *RequestHttp) GetRecordCount(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetRecords)

	if param == nil {
		param = &structs.GetRecordsReqParam{}
	}

	param.Limit = 1
	param.NeedFilterUserPermission = false
	param.IgnoreBackLookupField = false

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}
	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathGetRecords(namespace, objectAPIName), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	return gjson.GetBytes(data, "total").Int(), nil
}

func (r *RequestHttp) GetRecordCountV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetRecordsV2)

	if param == nil {
		param = &structs.GetRecordsReqParamV2{}
	}

	param.Limit = 1

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}
	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathGetRecordsV2(namespace, objectAPIName), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	return gjson.GetBytes(data, "total").Int(), nil
}

func (r *RequestHttp) CreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.CreateRecord)

	body := map[string]interface{}{
		"data":     record,
		"operator": cUtils.GetUserIDFromCtx(ctx),
		"task_id":  cUtils.GetTriggerTaskIDFromCtx(ctx),
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}
	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathCreateRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	id := gjson.GetBytes(data, "record_id").Int()
	if id == 0 {
		return nil, cExceptions.InternalError("id is empty, data: %s", string(data))
	}
	return &structs.RecordID{ID: id}, nil
}

func (r *RequestHttp) CreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.CreateRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathCreateRecordV2(namespace, objectAPIName), nil, record, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	id := gjson.GetBytes(data, "_id").Int()
	if id == 0 {
		return nil, cExceptions.InternalError("id is empty, data: %s", string(data))
	}
	return &structs.RecordID{ID: id}, nil
}

func (r *RequestHttp) BatchCreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchCreateRecord)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	body := map[string]interface{}{
		"data":               records,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
		"set_system_mod":     2,
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchCreateRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	result := structs.BatchCreateRecord{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("BatchCreateRecord failed, err: %v", err)
	}

	return result.RecordIDs, nil
}

func (r *RequestHttp) BatchCreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchCreateRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	body := map[string]interface{}{
		"records": records,
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchCreateRecordV2(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	result := structs.BatchCreateRecordV2{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("BatchCreateRecordV2 failed, err: %v", err)
	}

	return result.RecordIDs, nil
}

func (r *RequestHttp) BatchCreateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchCreateRecordAsync)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}

	body := map[string]interface{}{
		"data":               records,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchCreateRecordAsync(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	result := structs.AsyncTaskResult{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return 0, cExceptions.InternalError("CreateRecord failed, err: %v", err)
	}

	return result.TaskID, nil
}

func (r *RequestHttp) UpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UpdateRecord)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"record_id": recordID,
		"data":      record,
		"operator":  cUtils.GetUserIDFromCtx(ctx),
		"task_id":   cUtils.GetTriggerTaskIDFromCtx(ctx),
	}
	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathUpdateRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) UpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UpdateRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	_, err = errorWrapper(getOpenapiClient().PatchJson(ctx, GetPathUpdateRecordV2(namespace, objectAPIName, recordID), nil, record, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchUpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchUpdateRecord)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"data":               records,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
		"set_system_mod":     2,
	}
	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchUpdateRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchUpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchUpdateRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	newRecordMap := map[int64]map[string]interface{}{}
	err = cUtils.Decode(records, &newRecordMap)
	if err != nil {
		return cExceptions.InvalidParamError("Decode records failed, err: %+v", err)
	}

	var newRecordList []interface{}
	for id, record := range newRecordMap {
		if record == nil {
			continue
		}

		newRecordMap[id]["_id"] = id
		newRecordList = append(newRecordList, newRecordMap[id])
	}

	body := map[string]interface{}{
		"records": newRecordList,
	}
	_, err = errorWrapper(getOpenapiClient().PatchJson(ctx, GetPathBatchUpdateRecordV2(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchUpdateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchUpdateRecordAsync)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}

	body := map[string]interface{}{
		"data":               records,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
	}
	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchUpdateRecordAsync(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	result := structs.AsyncTaskResult{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return 0, cExceptions.InternalError("CreateRecord failed, err: %v", err)
	}
	return result.TaskID, nil
}

func (r *RequestHttp) DeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.DeleteRecord)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"record_id": recordID,
		"operator":  cUtils.GetUserIDFromCtx(ctx),
		"task_id":   cUtils.GetTriggerTaskIDFromCtx(ctx),
	}
	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathDeleteRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) DeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.DeleteRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	_, err = errorWrapper(getOpenapiClient().DeleteJson(ctx, GetPathDeleteRecordV2(namespace, objectAPIName, recordID), nil, map[string]interface{}{}, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchDeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchDeleteRecord)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"record_id_list":     recordIDs,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
		"set_system_mod":     2,
	}
	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchDeleteRecord(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchDeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchDeleteRecordV2)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"_ids": recordIDs,
	}
	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchDeleteRecordV2(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	return err
}

func (r *RequestHttp) BatchDeleteRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.BatchDeleteRecordAsync)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}

	body := map[string]interface{}{
		"record_id_list":     recordIDs,
		"operator":           cUtils.GetUserIDFromCtx(ctx),
		"automation_task_id": cUtils.GetTriggerTaskIDFromCtx(ctx),
	}
	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathBatchDeleteRecordAsync(namespace, objectAPIName), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	result := structs.AsyncTaskResult{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return 0, cExceptions.InternalError("BatchDeleteRecordAsync failed, err: %v", err)
	}
	return result.TaskID, nil
}

func (r *RequestHttp) Transaction(ctx context.Context, appCtx *structs.AppCtx, placeholders map[string]int64, operations []*structs.TransactionOperation) (map[string]int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.ModifyRecordsWithTransaction)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	body := map[string]interface{}{
		"placeholders":   placeholders,
		"operations":     operations,
		"operatorId":     cUtils.GetUserIDFromCtx(ctx),
		"taskId":         cUtils.GetTriggerTaskIDFromCtx(ctx),
		"setSystemField": 1,
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathTransaction(namespace), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result struct {
		Placeholders map[string]int64 `json:"placeholders"`
	}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("Oql failed, err: %v", err)
	}

	return result.Placeholders, nil
}

func (r *RequestHttp) Oql(ctx context.Context, appCtx *structs.AppCtx, oql string, args interface{}, namedArgs map[string]interface{}, records interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.Oql)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	body := map[string]interface{}{
		"query":     oql,
		"args":      args,
		"namedArgs": namedArgs,
		"compat":    true,
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathOql(namespace), nil, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return err
	}

	err = cUtils.JsonUnmarshalBytes([]byte(gjson.GetBytes(data, "rows").Raw), &records)
	if err != nil {
		return cExceptions.InternalError("Oql failed, err: %v", err)
	}

	return err
}

func (r *RequestHttp) DownloadFile(ctx context.Context, appCtx *structs.AppCtx, fileID string) ([]byte, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.DownloadAttachmentV2)

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetPathDownloadFileV2(fileID), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *RequestHttp) DownloadAvatar(ctx context.Context, appCtx *structs.AppCtx, imageID string) ([]byte, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.DownloadAvatar)

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetPathDownloadAvatar(imageID), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *RequestHttp) UploadFile(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader, expireSeconds time.Duration) (*structs.Attachment, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UploadAttachment)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if expireSeconds < time.Second && expireSeconds != 0 {
		return nil, cExceptions.InvalidParamError("expire time should be larger than one second or zero.")
	}

	err := writer.WriteField("expireSeconds", strconv.FormatInt(int64(expireSeconds.Seconds()), 10))
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile create expireSeconds failed, err: %v", err)
	}

	err = writer.WriteField("ignoreUserId", "true")
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile create ignoreUserId failed, err: %v", err)
	}

	form, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile create file failed, err: %v", err)
	}

	_, err = io.Copy(form, fileReader)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	headers := map[string][]string{
		cConstants.HttpHeaderKeyContentType: {writer.FormDataContentType()},
	}

	data, err := errorWrapper(getOpenapiClient().PostFormData(ctx, PathUploadFile, headers, payload, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result structs.Attachment
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFile failed, err: %v", err)
	}

	return &result, nil
}

func (r *RequestHttp) UploadFileV2(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.FileUploadResult, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UploadAttachmentV2)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	form, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		return nil, cExceptions.InternalError("UploadFileV2 create file failed, err: %v", err)
	}

	_, err = io.Copy(form, fileReader)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFileV2 failed, err: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, cExceptions.InternalError("UploadFileV2 failed, err: %v", err)
	}

	headers := map[string][]string{
		cConstants.HttpHeaderKeyContentType: {writer.FormDataContentType()},
	}

	data, err := errorWrapper(getOpenapiClient().PostFormData(ctx, PathUploadFileV2, headers, payload, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result structs.FileUploadResult
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("UploadFileV2 failed, err: %v", err)
	}

	return &result, nil
}

func (r *RequestHttp) UploadAvatar(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.Avatar, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UploadAvatar)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	form, err := writer.CreateFormFile("image", filepath.Base(fileName))
	if err != nil {
		return nil, cExceptions.InternalError("UploadAvatar create file failed, err: %v", err)
	}

	_, err = io.Copy(form, fileReader)
	if err != nil {
		return nil, cExceptions.InternalError("UploadAvatar failed, err: %v", err)
	}

	err = writer.Close()
	if err != nil {
		return nil, cExceptions.InternalError("UploadAvatar failed, err: %v", err)
	}

	headers := map[string][]string{
		cConstants.HttpHeaderKeyContentType: {writer.FormDataContentType()},
	}

	data, err := errorWrapper(getOpenapiClient().PostFormData(ctx, PathUploadAvatar, headers, payload, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	var result structs.Avatar
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return nil, cExceptions.InternalError("UploadAvatar failed, err: %v", err)
	}

	return &result, nil
}

func (r *RequestHttp) CreateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) (int64, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.CreateMessage)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathCreateMessage(namespace), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	// 不同构，openapi 返回 taskID，innerapi 返回 {taskID: 123}
	result := structs.AsyncTaskResult{}
	err = cUtils.JsonUnmarshalBytes(data, &result)
	if err != nil {
		return 0, cExceptions.InternalError("CreateMessage failed, err: %v", err)
	}

	return result.TaskID, nil
}

func (r *RequestHttp) UpdateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.UpdateMessage)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}

	_, err = errorWrapper(getOpenapiClient().PostJson(ctx, GetPathUpdateMessage(namespace), nil, param, cHttp.AppTokenMiddleware))
	if err != nil {
		return err
	}

	return nil
}

// GetGlobalConfig
func (r *RequestHttp) GetGlobalConfig(ctx context.Context, appCtx *structs.AppCtx, key string) (string, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetAllGlobalConfigs)

	pageSize, offset := 100, 0
	page := map[string]interface{}{
		"offset": 0,
		"limit":  pageSize,
	}
	body := map[string]interface{}{
		"biz_type": "GlobalVariables",
		"used_by":  "UsedBySystem",
		"filter":   page,
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return "", err
	}

	for i := 0; ; i++ {
		page["offset"] = offset * i
		data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathGetAllGlobalConfig(namespace), nil, body, cHttp.AppTokenMiddleware))
		if err != nil {
			return "", err
		}

		keyToValue := structs.GlobalConfigResult{}
		err = cUtils.JsonUnmarshalBytes(data, &keyToValue)
		if err != nil {
			return "", cExceptions.InternalError("GetAllGlobalConfig failed, err: %v", err)
		}

		for _, c := range keyToValue.Configs {
			if c.Key == key {
				return c.Value, nil
			}
		}

		if len(keyToValue.Configs) < pageSize {
			break
		}
	}

	return "", cExceptions.InvalidParamError("The global config (%s) does not exist", key)
}

func (r *RequestHttp) GetFields(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string) (result *structs.ObjFields, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetFields)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetPathGetFields(namespace, objectAPIName), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	if err := cUtils.JsonUnmarshalBytes(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RequestHttp) GetField(ctx context.Context, appCtx *structs.AppCtx, objectAPIName, fieldAPIName string) (result *structs.Field, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetField)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}

	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetPathGetField(namespace, objectAPIName, fieldAPIName), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}

	if err := cUtils.JsonUnmarshalBytes(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RequestHttp) MGetUserSettings(ctx context.Context, appCtx *structs.AppCtx, userIDList []int64) (result []*structs.Locale, err error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.MGetUserSettings)

	data, err := errorWrapper(getOpenapiClient().PostJson(ctx, GetPathMGetUserSettings(), nil, map[string]interface{}{
		"userIDs": userIDList,
	}, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RequestHttp) InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.InvokeFuncWithAuth)

	body, err := reqCommon.BuildInvokeParamsObj(ctx, apiName, params)
	if err != nil {
		return err
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}
	headers := map[string][]string{
		cConstants.HttpHeaderKeyUser: {strconv.FormatInt(cUtils.GetUserIDFromCtx(ctx), 10)},
	}

	respBody, extra, err := getOpenapiClient().PostJson(ctx, GetPathInvokeFunctionWithAuth(namespace, apiName), headers, body, cHttp.AppTokenMiddleware)
	data, err := errorWrapper(respBody, extra, err)
	if err != nil {
		return err
	}

	var resp struct {
		Result string `json:"result"`
	}

	logid := cUtils.GetLogIDFromExtra(extra)
	if err := cUtils.JsonUnmarshalBytes(data, &resp); err != nil {
		return cExceptions.InternalError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
	}

	code := gjson.GetBytes([]byte(resp.Result), "code").String()
	if code != "0" {
		msg := gjson.GetBytes([]byte(resp.Result), "msg").String()
		return cExceptions.InvalidParamError("%v ([%v] %v)", msg, code, logid)
	}

	dataRaw := gjson.GetBytes([]byte(resp.Result), "data").Raw
	if len(dataRaw) > 0 {
		if err := cUtils.JsonUnmarshalBytes([]byte(dataRaw), result); err != nil {
			return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v", err)
		}
	}

	return nil
}

func (r *RequestHttp) InvokeFunctionAsync(ctx context.Context, appCtx *structs.AppCtx, apiName string, params map[string]interface{}) (int64, error) {
	body, err := reqCommon.BuildInvokeParamsStr(ctx, apiName, params)
	if err != nil {
		return 0, err
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return 0, err
	}
	tenantName, err := utils.GetTenantName(ctx, appCtx)
	if err != nil {
		return 0, err
	}
	headers := map[string][]string{
		cConstants.HttpHeaderKeyTenant: {tenantName},
		cConstants.HttpHeaderKeyUser:   {strconv.FormatInt(cUtils.GetUserIDFromCtx(ctx), 10)},
	}

	data, err := errorWrapper(getOpenapiClient().PostJson(utils.SetAppConfToCtx(ctx, appCtx), GetPathInvokeFunctionAsync(namespace), headers, body, cHttp.AppTokenMiddleware))
	if err != nil {
		return 0, err
	}

	return gjson.GetBytes(data, "task_id").Int(), nil
}

func (r *RequestHttp) GetTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.TenantAccessToken, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetIntegrationTenantAccessToken)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}
	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetTenantAccessTokenPath(namespace, apiName), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}
	var inRes intern.TenantAccessToken
	err = cUtils.JsonUnmarshalBytes(data, &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[GetTenantAccessToken] failed, err: %v", err)
	}
	return &structs.TenantAccessToken{
		Expire:            inRes.Expire,
		TenantAccessToken: inRes.TenantAccessToken,
		AppID:             inRes.AppID,
	}, nil
}

func (r *RequestHttp) GetAppAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.AppAccessToken, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetIntegrationAppAccessToken)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}
	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetAppAccessTokenPath(namespace, apiName), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}
	var inRes intern.AppAccessToken
	err = cUtils.JsonUnmarshalBytes(data, &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[GetTenantAccessToken] failed, err: %v", err)
	}
	return &structs.AppAccessToken{
		Expire:         inRes.Expire,
		AppAccessToken: inRes.AppAccessToken,
		AppID:          inRes.AppID,
	}, nil
}

func (r *RequestHttp) GetDefaultTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.TenantAccessToken, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetDefaultIntegrationTenantAccessToken)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}
	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetDefaultTenantAccessTokenPath(namespace), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}
	var inRes intern.TenantAccessToken
	err = cUtils.JsonUnmarshalBytes(data, &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[GetTenantAccessToken] failed, err: %v", err)
	}
	return &structs.TenantAccessToken{
		Expire:            inRes.Expire,
		TenantAccessToken: inRes.TenantAccessToken,
		AppID:             inRes.AppID,
	}, nil
}

func (r *RequestHttp) GetDefaultAppAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.AppAccessToken, error) {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.GetDefaultIntegrationAppAccessToken)

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return nil, err
	}
	data, err := errorWrapper(getOpenapiClient().Get(ctx, GetDefaultAppAccessTokenPath(namespace), nil, cHttp.AppTokenMiddleware))
	if err != nil {
		return nil, err
	}
	var inRes intern.AppAccessToken
	err = cUtils.JsonUnmarshalBytes(data, &inRes)
	if err != nil {
		return nil, cExceptions.InvalidParamError("[GetTenantAccessToken] failed, err: %v", err)
	}
	return &structs.AppAccessToken{
		Expire:         inRes.Expire,
		AppAccessToken: inRes.AppAccessToken,
		AppID:          inRes.AppID,
	}, nil
}
