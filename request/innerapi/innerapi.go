package innerapi

//type RequestRpc struct{}
//
//func (r *RequestRpc) GetTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.TenantAccessToken, error) {
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.ExecuteFlow)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//	req := tenant_manager.NewGetIntegrationTenantAccessTokenRequest()
//	req.APIName = apiName
//
//	defer cancel()
//	resp, err := cli.GetIntegrationTenantAccessToken(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return &structs.TenantAccessToken{
//		Expire:            *resp.Expire,
//		TenantAccessToken: *resp.TenantAccessToken,
//		AppID:             *resp.AppID,
//	}, nil
//}
//
//func (r *RequestRpc) GetAppAccessToken(ctx context.Context, appCtx *structs.AppCtx, apiName string) (*structs.AppAccessToken, error) {
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.ExecuteFlow)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//	req := tenant_manager.NewGetIntegrationAppAccessTokenRequest()
//	req.APIName = apiName
//
//	defer cancel()
//	resp, err := cli.GetIntegrationAppAccessToken(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return &structs.AppAccessToken{
//		Expire:         *resp.Expire,
//		AppAccessToken: *resp.AppAccessToken,
//		AppID:          *resp.AppID,
//	}, nil
//}
//
//func (r *RequestRpc) GetDefaultTenantAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.TenantAccessToken, error) {
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.ExecuteFlow)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//	req := tenant_manager.NewGetDefaultIntegrationTenantAccessTokenRequest()
//
//	defer cancel()
//	resp, err := cli.GetDefaultIntegrationTenantAccessToken(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return &structs.TenantAccessToken{
//		Expire:            *resp.Expire,
//		TenantAccessToken: *resp.TenantAccessToken,
//		AppID:             *resp.AppID,
//	}, nil
//}
//
//func (r *RequestRpc) GetDefaultAppAccessToken(ctx context.Context, appCtx *structs.AppCtx) (*structs.AppAccessToken, error) {
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.ExecuteFlow)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//	req := tenant_manager.NewGetDefaultIntegrationAppAccessTokenRequest()
//
//	defer cancel()
//	resp, err := cli.GetDefaultIntegrationAppAccessToken(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return &structs.AppAccessToken{
//		Expire:         *resp.Expire,
//		AppAccessToken: *resp.AppAccessToken,
//		AppID:          *resp.AppID,
//	}, nil
//}
//
//func (r *RequestRpc) Execute(ctx context.Context, appCtx *structs.AppCtx, APIName string, options *structs.ExecuteOptions) (invokeResult *structs.FlowExecuteResult, err error) {
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.ExecuteFlow)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	req := workflow.NewExecuteRequest()
//	req.NameSpace = namespace
//	req.FlowAPIName = APIName
//	req.Operator = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.LoopMasks = cUtils.GetLoopMaskFromCtx(ctx)
//
//	varStr, _ := cUtils.JsonMarshalBytes(utils.ParseMapToFlowVariable(options.Params))
//	req.Variables = cUtils.StringPtr(string(varStr))
//
//	defer cancel()
//	resp, err := cli.Execute(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	var outParams intern.ExecuteFlowVariables
//	if resp != nil && resp.OutParams != nil && *resp.OutParams != "" {
//		if err := cUtils.JsonUnmarshalBytes([]byte(*resp.OutParams), &outParams); err != nil {
//			return nil, cExceptions.InternalError("[Execute] outParams decode err: %v", err)
//		}
//	}
//	return &structs.FlowExecuteResult{
//		ExecutionID: resp.ExecutionID,
//		Status:      structs.ExecutionStatus(resp.Status),
//		Data:        utils.ParseFlowVariableToMap(outParams),
//		ErrCode:     resp.Code,
//		ErrMsg:      resp.ErrorMsg,
//	}, nil
//}
//
//func (r *RequestRpc) RevokeExecution(ctx context.Context, appCtx *structs.AppCtx, executionID int64, options *structs.RevokeOptions) (err error) {
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.RevokeExecution)
//	if err != nil {
//		return err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	req := workflow.NewRevokeExecutionRequest()
//	req.NameSpace = namespace
//	req.ExecutionID = executionID
//	req.Operator = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//
//	if options != nil {
//		req.Reason = &workflow.Multilingual{}
//		if options.Reason.Zh != "" {
//			req.Reason.Zh = &options.Reason.Zh
//		}
//		if options.Reason.En != "" {
//			req.Reason.En = &options.Reason.En
//		}
//	}
//
//	defer cancel()
//	resp, err := cli.RevokeExecution(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//	return
//}
//
//func (r *RequestRpc) GetExecutionInfo(ctx context.Context, appCtx *structs.AppCtx, executionID int64) (*structs.ExecutionInfo, error) {
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetExecutionInfo)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	req := workflow.NewGetExecutionInfoRequest()
//	req.NameSpace = namespace
//	req.ExecutionID = executionID
//	req.Operator = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//
//	defer cancel()
//	resp, err := cli.GetExecutionInfo(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	if resp.ExecutionInfo == nil {
//		return nil, cExceptions.InternalError("[GetExecutionInfo] ExecutionInfo is empty")
//	}
//	var outParams intern.ExecuteFlowVariables
//	if resp != nil && resp.ExecutionInfo.OutParams != nil && *resp.ExecutionInfo.OutParams != "" {
//		if err := cUtils.JsonUnmarshalBytes([]byte(*resp.ExecutionInfo.OutParams), &outParams); err != nil {
//			return nil, cExceptions.InternalError("[GetExecutionInfo] outParams decode err: %v", err)
//		}
//	}
//	return &structs.ExecutionInfo{
//		Status:  structs.ExecutionStatus(resp.ExecutionInfo.Status),
//		Data:    utils.ParseFlowVariableToMap(outParams),
//		ErrCode: resp.ExecutionInfo.Code,
//		ErrMsg:  resp.ExecutionInfo.ErrorMsg,
//	}, nil
//}
//
//func (r *RequestRpc) GetExecutionUserTaskInfo(ctx context.Context, appCtx *structs.AppCtx, executionID int64) (taskInfoList []*structs.TaskInfo, err error) {
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetExecutionUserTaskInfo)
//	if err != nil {
//		return nil, err
//	}
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	req := workflow.NewGetExecutionUserTaskInfoRequest()
//	req.NameSpace = namespace
//	req.ExecutionID = executionID
//	req.Operator = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//
//	defer cancel()
//	resp, err := cli.GetExecutionUserTaskInfo(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	if err := mapstructure.Decode(resp.GetTaskList(), &taskInfoList); err != nil {
//		return nil, cExceptions.InternalError("[GetExecutionUserTaskInfo] result decode failed: %v", err)
//	}
//
//	return taskInfoList, nil
//}
//
//func (r *RequestRpc) GetTenantInfo(ctx context.Context, appCtx *structs.AppCtx) (*cStructs.Tenant, error) {
//	ctx, _, _, err := r.pre(ctx, appCtx, cConstants.GetAppToken)
//	if err != nil {
//		return nil, err
//	}
//	return appCtx.Credential.GetTenantInfo(ctx)
//}
//
//func (r *RequestRpc) pre(ctx context.Context, appCtx *structs.AppCtx, method string) (context.Context, context.CancelFunc, string, error) {
//	var err error
//	ctx, err = cHttp.RebuildRpcCtx(utils.SetCtx(ctx, appCtx, method))
//	if err != nil {
//		return nil, nil, "", err
//	}
//
//	namespace, err := utils.GetNamespace(ctx, appCtx)
//	if err != nil {
//		return nil, nil, "", err
//	}
//
//	if namespace == "" {
//		return nil, nil, "", cExceptions.InternalError("namespace is empty")
//	}
//
//	var cancel context.CancelFunc
//	ctx, cancel = cHttp.GetTimeoutCtx(ctx)
//	return ctx, cancel, namespace, nil
//}
//
//func (r *RequestRpc) post(ctx context.Context, err error, baseResp *base.BaseResp, baseReq *base.Base) error {
//	var logid string
//	if baseReq != nil {
//		logid = baseReq.LogID
//	}
//
//	if err != nil {
//		return cExceptions.InternalError("Call InnerAPI failed: %+v, logid: %s", err, logid)
//	}
//
//	if baseResp == nil {
//		return cExceptions.InternalError("Call InnerAPI resp is empty, logid: %s", logid)
//	}
//
//	if baseResp.KStatusCode != "" {
//		msg := baseResp.KStatusMessage
//		if baseResp.StatusMessage != "" {
//			msg = baseResp.StatusMessage
//		}
//		return cExceptions.NewErrWithCodeV2(baseResp.KStatusCode, msg, logid)
//	}
//	return nil
//}
//
//func (r *RequestRpc) getRecordsRequest(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam) (*string, [][]string, error) {
//	if param == nil {
//		param = &structs.GetRecordsReqParam{}
//	}
//	if param.Limit <= 0 {
//		param.Limit = constants.PageLimitDefault
//	}
//	param.NeedFilterUserPermission = false
//	param.IgnoreBackLookupField = false
//
//	criterionBytes, err := cUtils.JsonMarshalBytes(param.Criterion)
//	if err != nil {
//		return nil, nil, cExceptions.InternalError("Marshal criterion failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchGetRecordWithCriterionRequestV3()
//	req.ObjectAPIAlias = objectAPIName
//	req.Criterion = string(criterionBytes)
//	req.Limit = param.Limit
//	req.Offset = param.Offset
//	req.FieldAPIAPIAliases = param.FieldApiNames
//	req.NeedTotalCount = cUtils.BoolPtr(false)
//	req.NeedFilterUserPermission = cUtils.BoolPtr(false)
//	req.IgnoreBackLookupField = false
//	if param.FuzzySearch != nil {
//		req.FuzzySearch = &metadata.FuzzySearch{
//			Keyword:       param.FuzzySearch.Keyword,
//			FieldAPINames: param.FuzzySearch.FieldAPINames,
//		}
//	}
//	authFieldType := unauth_permission.ProcessAuthFieldType_SliceResult
//	req.ProcessAuthFieldType = &authFieldType
//
//	for _, order := range param.Order {
//		req.Order = append(req.Order, &common.Sort{
//			FieldAPIName: order.Field,
//			Type:         order.Type,
//			Direction:    order.Direction,
//		})
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetRecords)
//	if err != nil {
//		return nil, nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	resp, err := cli.BatchGetRecordWithCriterionV3(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, nil, err
//	}
//
//	if resp.UnauthPermissionInfo != nil && len(resp.UnauthPermissionInfo.UnauthFieldSlice) > 0 {
//		return resp.DataList, resp.UnauthPermissionInfo.UnauthFieldSlice, nil
//	}
//	return resp.DataList, nil, nil
//}
//
//func (r *RequestRpc) GetRecords(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam, records interface{}) ([][]string, error) {
//	recordList, unauthFields, err := r.getRecordsRequest(ctx, appCtx, objectAPIName, param)
//	if err != nil {
//		return nil, err
//	}
//
//	if recordList == nil || len(*recordList) == 0 {
//		return nil, nil
//	}
//
//	_, ok1 := records.(*[]std_record.Record)
//	_, ok2 := records.(*[]*std_record.Record)
//	if ok1 || ok2 {
//		var newRecords []map[string]interface{}
//		err = cUtils.JsonUnmarshalBytes([]byte(*recordList), &newRecords)
//		if err != nil {
//			return nil, cExceptions.InvalidParamError("Unmarshal DataList failed: %+v", err)
//		}
//		cHttp.AppendUnauthFieldRecordList(ctx, objectAPIName, newRecords, unauthFields)
//
//		if len(unauthFields) > 0 && len(unauthFields) != len(newRecords) {
//			return nil, cExceptions.InternalError("len(records) %d != len(unauthFields) %d", len(newRecords), len(unauthFields))
//		}
//
//		rv := reflect.ValueOf(records).Elem()
//		arr := make([]reflect.Value, len(newRecords), len(newRecords))
//		for i, record := range newRecords {
//			v := std_record.Record{
//				Record: record,
//			}
//			if len(unauthFields) > 0 {
//				v.UnauthFields = unauthFields[i]
//			}
//			if ok1 {
//				arr[i] = reflect.ValueOf(v)
//			} else {
//				arr[i] = reflect.ValueOf(&v)
//			}
//		}
//
//		rv.Set(reflect.Append(rv, arr...))
//	} else {
//		err = cUtils.JsonUnmarshalBytes([]byte(*recordList), &records)
//		if err != nil {
//			return nil, cExceptions.InvalidParamError("Unmarshal DataList failed: %+v", err)
//		}
//		cHttp.AppendUnauthFieldRecordList(ctx, objectAPIName, records, unauthFields)
//	}
//
//	return unauthFields, nil
//}
//
//func (r *RequestRpc) getRecordsV2Request(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2) (string, [][]string, error) {
//	if param == nil {
//		param = &structs.GetRecordsReqParamV2{}
//	}
//	if param.Limit <= 0 {
//		param.Limit = constants.PageLimitDefault
//	}
//
//	filterBytes, err := cUtils.JsonMarshalBytes(param.Filter)
//	if err != nil {
//		return "", nil, cExceptions.InternalError("Marshal criterion failed, err: %+v", err)
//	}
//
//	req := datax.NewGetRecordListRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Filter = cUtils.StringPtr(string(filterBytes))
//	req.Limit = cUtils.Int64Ptr(param.Limit)
//	req.Offset = cUtils.Int64Ptr(param.Offset)
//	req.FiledAPIAlias = param.Fields
//	req.NeedTotalCount = cUtils.BoolPtr(false)
//	authFieldType := unauth_permission.ProcessAuthFieldType_SliceResult
//	req.ProcessAuthFieldType = &authFieldType
//
//	for _, order := range param.Sort {
//		req.Sort = append(req.Sort, &common.AliasSort{
//			FieldApiAlias: order.Field,
//			Type:          order.Type,
//			Direction:     order.Direction,
//		})
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetRecordsV2)
//	if err != nil {
//		return "", nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return "", nil, err
//	}
//
//	resp, err := cli.GetRecordList(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return "", nil, err
//	}
//
//	if resp.UnauthPermissionInfo != nil {
//		return resp.Data, resp.UnauthPermissionInfo.UnauthFieldSlice, nil
//	}
//	return resp.Data, nil, nil
//}
//
//func (r *RequestRpc) GetRecordsV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2, records interface{}) ([][]string, error) {
//	recordList, unauthFields, err := r.getRecordsV2Request(ctx, appCtx, objectAPIName, param)
//	if err != nil {
//		return nil, err
//	}
//
//	_, ok1 := records.(*[]std_record.Record)
//	_, ok2 := records.(*[]*std_record.Record)
//	if ok1 || ok2 {
//		var newRecords []map[string]interface{}
//		err = cUtils.JsonUnmarshalBytes([]byte(recordList), &newRecords)
//		if err != nil {
//			return nil, cExceptions.InvalidParamError("Unmarshal DataList failed: %+v", err)
//		}
//		cHttp.AppendUnauthFieldRecordList(ctx, objectAPIName, recordList, unauthFields)
//
//		if len(unauthFields) > 0 && len(unauthFields) != len(newRecords) {
//			return nil, cExceptions.InternalError("len(records) %d != len(unauthFields) %d", len(newRecords), len(unauthFields))
//		}
//
//		rv := reflect.ValueOf(records).Elem()
//		arr := make([]reflect.Value, len(newRecords), len(newRecords))
//		for i, record := range newRecords {
//			v := std_record.Record{
//				Record: record,
//			}
//			if len(unauthFields) > 0 {
//				v.UnauthFields = unauthFields[i]
//			}
//			if ok1 {
//				arr[i] = reflect.ValueOf(v)
//			} else {
//				arr[i] = reflect.ValueOf(&v)
//			}
//		}
//
//		rv.Set(reflect.Append(rv, arr...))
//	} else {
//		err = cUtils.JsonUnmarshalBytes([]byte(recordList), records)
//		if err != nil {
//			return nil, cExceptions.InvalidParamError("Unmarshal DataList failed: %+v", err)
//		}
//		cHttp.AppendUnauthFieldRecordList(ctx, objectAPIName, records, unauthFields)
//	}
//
//	return unauthFields, nil
//}
//
//func (r *RequestRpc) GetRecordCount(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParam) (int64, error) {
//	if param == nil {
//		param = &structs.GetRecordsReqParam{}
//	}
//	param.Limit = 1
//
//	criterionBytes, err := cUtils.JsonMarshalBytes(param.Criterion)
//	if err != nil {
//		return 0, cExceptions.InternalError("Marshal criterion failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchGetRecordWithCriterionRequestV3()
//	req.ObjectAPIAlias = objectAPIName
//	req.Criterion = string(criterionBytes)
//	req.Limit = param.Limit
//	req.Offset = param.Offset
//	req.FieldAPIAPIAliases = param.FieldApiNames
//	req.NeedTotalCount = cUtils.BoolPtr(true)
//	req.NeedFilterUserPermission = cUtils.BoolPtr(false)
//	req.IgnoreBackLookupField = false
//	for _, order := range param.Order {
//		req.Order = append(req.Order, &common.Sort{
//			FieldAPIName: order.Field,
//			Type:         order.Type,
//			Direction:    order.Direction,
//		})
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetRecords)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.BatchGetRecordWithCriterionV3(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.GetTotal(), nil
//}
//
//func (r *RequestRpc) GetRecordCountV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, param *structs.GetRecordsReqParamV2) (int64, error) {
//	if param == nil {
//		param = &structs.GetRecordsReqParamV2{}
//	}
//	param.Limit = 1
//
//	filterBytes, err := cUtils.JsonMarshalBytes(param.Filter)
//	if err != nil {
//		return 0, cExceptions.InternalError("Marshal criterion failed, err: %+v", err)
//	}
//
//	req := datax.NewGetRecordListRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Filter = cUtils.StringPtr(string(filterBytes))
//	req.Limit = cUtils.Int64Ptr(param.Limit)
//	req.Offset = cUtils.Int64Ptr(param.Offset)
//	req.FiledAPIAlias = param.Fields
//	req.NeedTotalCount = cUtils.BoolPtr(true)
//	for _, order := range param.Sort {
//		req.Sort = append(req.Sort, &common.AliasSort{
//			FieldApiAlias: order.Field,
//			Type:          order.Type,
//			Direction:     order.Direction,
//		})
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetRecordsV2)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.GetRecordList(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.GetTotal(), nil
//}
//
//func (r *RequestRpc) CreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error) {
//	newRecord := std_record.ConvertStdRecord(record)
//	recordBytes, err := cUtils.JsonMarshalBytes([]interface{}{newRecord})
//	if err != nil {
//		return nil, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchCreateRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.CreateRecord)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchCreateRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	return &structs.RecordID{
//		ID: resp.GetRecordIDs()[0],
//	}, nil
//}
//
//func (r *RequestRpc) CreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, record interface{}) (*structs.RecordID, error) {
//	newRecord := std_record.ConvertStdRecord(record)
//	recordBytes, err := cUtils.JsonMarshalBytes(newRecord)
//	if err != nil {
//		return nil, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := datax.NewCreateRecordRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Record = recordBytes
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.CreateRecordV2)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.CreateRecord(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	return &structs.RecordID{
//		ID: resp.RecordID,
//	}, nil
//}
//
//func (r *RequestRpc) BatchCreateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error) {
//	newRecords := std_record.ConvertStdRecords(records)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecords)
//	if err != nil {
//		return nil, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchCreateRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordsBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchCreateRecord)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchCreateRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	return resp.RecordIDs, nil
//}
//
//func (r *RequestRpc) BatchCreateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) ([]int64, error) {
//	newRecords := std_record.ConvertStdRecords(records)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecords)
//	if err != nil {
//		return nil, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := datax.NewBatchCreateRecordsRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Records = string(recordsBytes)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchCreateRecordV2)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchCreateRecords(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	return resp.RecordIDs, nil
//}
//
//func (r *RequestRpc) BatchCreateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records interface{}) (int64, error) {
//	newRecords := std_record.ConvertStdRecords(records)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecords)
//	if err != nil {
//		return 0, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchCreateRecordByAsyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordsBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchCreateRecordAsync)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.BatchCreateRecordByAsync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.TaskID, nil
//}
//
//func (r *RequestRpc) UpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error {
//	newRecord := std_record.ConvertStdRecord(record)
//	recordsBytes, err := cUtils.JsonMarshalBytes(map[int64]interface{}{
//		recordID: newRecord,
//	})
//	if err != nil {
//		return cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchUpdateRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordsBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.UpdateRecord)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.BatchUpdateRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) UpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64, record interface{}) error {
//	newRecord := std_record.ConvertStdRecord(record)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecord)
//	if err != nil {
//		return cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := datax.NewUpdateRecordRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordID = recordID
//	req.Record = recordsBytes
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.UpdateRecordV2)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.UpdateRecord(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (r *RequestRpc) BatchUpdateRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) (*structs.BatchResult, error) {
//	newRecords := std_record.ConvertStdRecordsFromMap(records)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecords)
//	if err != nil {
//		return nil, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchUpdateRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordsBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchUpdateRecord)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchUpdateRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return reqCommon.GenBatchResultByRecords(records, resp.ErrMap), nil
//}
//
//func (r *RequestRpc) BatchUpdateRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) error {
//	newRecords := std_record.ConvertStdRecordsFromMap(records)
//	newRecordMap := map[int64]map[string]interface{}{}
//	err := cUtils.Decode(newRecords, &newRecordMap)
//	if err != nil {
//		return cExceptions.InvalidParamError("Decode records failed, err: %+v", err)
//	}
//
//	var newRecordList []interface{}
//	for id, record := range newRecordMap {
//		if record == nil {
//			continue
//		}
//
//		newRecordMap[id]["_id"] = id
//		newRecordList = append(newRecordList, newRecordMap[id])
//	}
//
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecordList)
//	if err != nil {
//		return cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := datax.NewBatchUpdateRecordsRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Records = string(recordsBytes)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchUpdateRecordV2)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.BatchUpdateRecords(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) BatchUpdateRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, records map[int64]interface{}) (int64, error) {
//	newRecords := std_record.ConvertStdRecordsFromMap(records)
//	recordsBytes, err := cUtils.JsonMarshalBytes(newRecords)
//	if err != nil {
//		return 0, cExceptions.InvalidParamError("Marshhal record failed, err: %+v", err)
//	}
//
//	req := metadata.NewBatchUpdateRecordByAsyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.Data = string(recordsBytes)
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchUpdateRecordAsync)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.BatchUpdateRecordByAsync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.TaskID, nil
//}
//
//func (r *RequestRpc) DeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error {
//	req := metadata.NewBatchDeleteRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordIDs = []int64{recordID}
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.DeleteRecord)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.BatchDeleteRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) DeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordID int64) error {
//	req := datax.NewDeleteRecordRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordID = recordID
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.DeleteRecordV2)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.DeleteRecord(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) BatchDeleteRecord(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) (*structs.BatchResult, error) {
//	req := metadata.NewBatchDeleteRecordBySyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordIDs = recordIDs
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.SetSystemModPtr(common.SetSystemMod_Other)
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchDeleteRecord)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchDeleteRecordBySync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return reqCommon.GenBatchResultByRecordIDs(recordIDs, resp.ErrMap), nil
//}
//
//func (r *RequestRpc) BatchDeleteRecordV2(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) error {
//	req := datax.NewBatchDeleteRecordsRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordIDs = recordIDs
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchDeleteRecordV2)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.BatchDeleteRecords(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) BatchDeleteRecordAsync(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string, recordIDs []int64) (int64, error) {
//	req := metadata.NewBatchDeleteRecordByAsyncRequest()
//	req.ObjectAPIAlias = objectAPIName
//	req.RecordIDList = recordIDs
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.AutomationTaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.BatchDeleteRecordAsync)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.BatchDeleteRecordByAsync(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.TaskID, nil
//}
//
//func (r *RequestRpc) oqlRequest(ctx context.Context, appCtx *structs.AppCtx, oql string, args interface{}, namedArgs map[string]interface{}) (rows string, unauthFields [][]string, err error) {
//	param, err := cUtils.JsonMarshalBytes(map[string]interface{}{
//		"query":        oql,
//		"args":         args,
//		"namedArgs":    namedArgs,
//		"compat":       true,
//		"unauthFields": true,
//	})
//	if err != nil {
//		return rows, nil, cExceptions.InternalError("Marshal failed, err: %+v", err)
//	}
//
//	req := metadata.NewQueryRecordsByOQLRequest()
//	req.QueryData = param
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.Oql)
//	if err != nil {
//		return rows, nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return rows, nil, err
//	}
//
//	resp, err := cli.QueryRecordsByOQL(ctx, req)
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return rows, nil, err
//	}
//
//	return resp.Rows, resp.UnauthFieldSlice, nil
//}
//
//func (r *RequestRpc) Oql(ctx context.Context, appCtx *structs.AppCtx, oql string, args interface{}, namedArgs map[string]interface{}, resultSet interface{}) (unauthFields [][]string, err error) {
//	recordList, unauthFields, err := r.oqlRequest(ctx, appCtx, oql, args, namedArgs)
//	if err != nil {
//		return nil, err
//	}
//
//	_, ok1 := resultSet.(*[]std_record.Record)
//	_, ok2 := resultSet.(*[]*std_record.Record)
//	if ok1 || ok2 {
//		var newRecords []map[string]interface{}
//		err = cUtils.JsonUnmarshalBytes([]byte(recordList), &newRecords)
//		if err != nil {
//			return nil, cExceptions.InvalidParamError("Unmarshal DataList failed: %+v", err)
//		}
//
//		if len(unauthFields) > 0 && len(unauthFields) != len(newRecords) {
//			return nil, cExceptions.InternalError("len(records) %d != len(unauthFields) %d", len(newRecords), len(unauthFields))
//		}
//
//		rv := reflect.ValueOf(resultSet).Elem()
//		arr := make([]reflect.Value, len(newRecords), len(newRecords))
//		for i, record := range newRecords {
//			v := std_record.Record{
//				Record: record,
//			}
//			if len(unauthFields) > 0 {
//				v.UnauthFields = unauthFields[i]
//			}
//			if ok1 {
//				arr[i] = reflect.ValueOf(v)
//			} else {
//				arr[i] = reflect.ValueOf(&v)
//			}
//		}
//
//		rv.Set(reflect.Append(rv, arr...))
//	} else {
//		err = cUtils.JsonUnmarshalBytes([]byte(recordList), &resultSet)
//		if err != nil {
//			return nil, cExceptions.InternalError("Oql failed, err: %v", err)
//		}
//	}
//	return unauthFields, nil
//}
//
//func (r *RequestRpc) Transaction(ctx context.Context, appCtx *structs.AppCtx, placeholders map[string]int64, operations []*structs.TransactionOperation) (map[string]int64, error) {
//	req := datax.NewModifyRecordsWithTXRequest()
//	req.Placeholders = placeholders
//	req.OperatorID = cUtils.Int64Ptr(cUtils.GetUserIDFromCtx(ctx))
//	req.TaskID = cUtils.Int64Ptr(cUtils.GetTriggerTaskIDFromCtx(ctx))
//	req.SetSystemField = common.CommitSetSystemModPtr(common.CommitSetSystemMod_SysFieldSet)
//
//	for _, operation := range operations {
//		req.Operations = append(req.Operations, &common.Operation{
//			OperationType: common.OperationType(operation.OperationType),
//			ObjectAPIName: operation.ObjectAPIName,
//			Input:         cUtils.StringPtr(operation.Input),
//		})
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.ModifyRecordsWithTransaction)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.ModifyRecordsWithTX(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	return resp.Placeholders, nil
//}
//
//func (r *RequestRpc) DownloadFile(ctx context.Context, appCtx *structs.AppCtx, fileID string) ([]byte, error) {
//	req := attachment.NewDownloadFileRequest()
//	req.FileID = fileID
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.DownloadAttachmentV2)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.DownloadFile(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return resp.Data, nil
//}
//
//func (r *RequestRpc) DownloadAvatar(ctx context.Context, appCtx *structs.AppCtx, imageID string) ([]byte, error) {
//	req := attachment.NewDownloadAvatarRequest()
//	req.ImageID = imageID
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.DownloadAvatar)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.DownloadAvatar(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return resp.Data, nil
//}
//
//func (r *RequestRpc) UploadFile(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader, expireSeconds time.Duration) (*structs.Attachment, error) {
//	req := attachment.NewUploadAttachmentForRPCRequest()
//	data, err := ioutil.ReadAll(fileReader)
//	if err != nil {
//		return nil, cExceptions.InternalError("ioutil.ReadAll failed, err: %+v", err)
//	}
//
//	if expireSeconds < time.Second && expireSeconds != 0 {
//		return nil, cExceptions.InvalidParamError("expire time should be larger than one second or zero.")
//	}
//
//	req.Data = data
//	req.Name = fileName
//	req.EncryptType = common.EncryptTypePtr(common.EncryptType_NoEncrypt)
//	req.ExpireSecond = cUtils.Int64Ptr(int64(expireSeconds.Seconds()))
//	req.PermissionConf = &common.PermissionConfig{IgnoreUserID: true}
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.UploadAttachment)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.UploadAttachmentForRPC(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	result := &structs.Attachment{
//		ID:    resp.ID,
//		Token: resp.Token,
//		Size:  resp.Size,
//	}
//	if resp.MimeType != nil {
//		result.MimeType = *resp.MimeType
//	}
//	if resp.Filename != nil {
//		result.Name = *resp.Filename
//	}
//	return result, nil
//}
//
//func (r *RequestRpc) UploadFileV2(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.FileUploadResult, error) {
//	req := attachment.NewUploadFileForRPCRequest()
//	data, err := ioutil.ReadAll(fileReader)
//	if err != nil {
//		return nil, cExceptions.InternalError("ioutil.ReadAll failed, err: %+v", err)
//	}
//
//	req.Data = data
//	req.Name = fileName
//	req.EncryptType = common.EncryptTypePtr(common.EncryptType_NoEncrypt)
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.UploadAttachmentV2)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.UploadFileForRPC(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	result := &structs.FileUploadResult{
//		FileID: resp.FileID,
//		Size:   int(resp.Size),
//	}
//	if resp.MimeType != nil {
//		result.Type = *resp.MimeType
//	}
//	if resp.Filename != nil {
//		result.Name = *resp.Filename
//	}
//	return result, nil
//}
//
//func (r *RequestRpc) UploadAvatar(ctx context.Context, appCtx *structs.AppCtx, fileName string, fileReader io.Reader) (*structs.Avatar, error) {
//	req := attachment.NewUploadAvatarForRPCRequest()
//	data, err := ioutil.ReadAll(fileReader)
//	if err != nil {
//		return nil, cExceptions.InternalError("ioutil.ReadAll failed, err: %+v", err)
//	}
//
//	req.Data = data
//	req.EncryptType = common.EncryptTypePtr(common.EncryptType_NoEncrypt)
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.UploadAvatar)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.UploadAvatarForRPC(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	result := &structs.Avatar{
//		ImageID:        resp.ImageID,
//		PreviewImageID: resp.PreviewImageId,
//	}
//	return result, nil
//}
//
//func (r *RequestRpc) CreateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) (int64, error) {
//	req := message.NewCreateNotifyTaskRequest()
//	err := cUtils.Decode(param, &req)
//	if err != nil {
//		return 0, cExceptions.InternalError("Decode req failed, err: %+v", err)
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.CreateMessage)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//	req.NameSpace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.CreateNotifyTask(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.TaskID, nil
//}
//
//func (r *RequestRpc) UpdateMessage(ctx context.Context, appCtx *structs.AppCtx, param map[string]interface{}) error {
//	req := message.NewUpdateNotifyTaskRequest()
//	err := cUtils.Decode(param, &req)
//	if err != nil {
//		return cExceptions.InternalError("Decode req failed, err: %+v", err)
//	}
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.UpdateMessage)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	req.NameSpace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.UpdateNotifyTask(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) GetGlobalConfig(ctx context.Context, appCtx *structs.AppCtx, key string) (string, error) {
//	pageSize, offset := int64(100), int64(0)
//	filter := &common.Filter{
//		Offset: cUtils.Int64Ptr(offset),
//		Limit:  cUtils.Int64Ptr(pageSize),
//	}
//	req := globalconfig.NewGetAllConfigRequest()
//	req.BizType = "GlobalVariables"
//	req.UsedBy = "UsedBySystem"
//	req.Filter = filter
//
//	for i := 0; ; i++ {
//		filter.Offset = cUtils.Int64Ptr(offset * int64(i))
//		configs, err := r.getConfigByPage(ctx, appCtx, req)
//		if err != nil {
//			return "", err
//		}
//
//		for _, config := range configs {
//			if config.Key == key {
//				return config.Value, nil
//			}
//		}
//
//		if int64(len(configs)) < pageSize {
//			break
//		}
//	}
//
//	return "", cExceptions.InvalidParamError("The global config (%s) does not exist", key)
//}
//
//func (r *RequestRpc) GetAllGlobalConfig(ctx context.Context, appCtx *structs.AppCtx) (map[string]string, error) {
//	pageSize, offset := int64(100), int64(0)
//	filter := &common.Filter{
//		Offset: cUtils.Int64Ptr(offset),
//		Limit:  cUtils.Int64Ptr(pageSize),
//	}
//	req := globalconfig.NewGetAllConfigRequest()
//	req.BizType = "GlobalVariables"
//	req.UsedBy = "UsedBySystem"
//	req.Filter = filter
//
//	keyToValue := map[string]string{}
//	for i := 0; ; i++ {
//		filter.Offset = cUtils.Int64Ptr(offset * int64(i))
//		configs, err := r.getConfigByPage(ctx, appCtx, req)
//		if err != nil {
//			return nil, err
//		}
//
//		for _, config := range configs {
//			keyToValue[config.Key] = config.Value
//		}
//
//		if int64(len(configs)) < pageSize {
//			break
//		}
//	}
//
//	return keyToValue, nil
//}
//
//func (r *RequestRpc) getConfigByPage(ctx context.Context, appCtx *structs.AppCtx, req *globalconfig.GetAllConfigRequest) ([]*common.Config, error) {
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetAllGlobalConfigs)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.NameSpace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.GetAllConfig(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//	return resp.Configs, nil
//}
//
//func (r *RequestRpc) GetFields(ctx context.Context, appCtx *structs.AppCtx, objectAPIName string) (result *structs.ObjFields, err error) {
//	req := metadata.NewGetObjectFieldsRequest()
//	req.ObjectAPIName = objectAPIName
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetFields)
//	if err != nil {
//		return nil, err
//	}
//	req.Namespace = namespace
//	defer cancel()
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.GetObjectFields(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	if err := cUtils.JsonUnmarshalBytes([]byte(resp.GetData()), &result); err != nil {
//		return nil, err
//	}
//	return result, nil
//
//}
//
//func (r *RequestRpc) GetField(ctx context.Context, appCtx *structs.AppCtx, objectAPIName, fieldAPIName string) (result *structs.Field, err error) {
//	req := metadata.NewGetObjectFieldRequest()
//	req.ObjectAPIName = objectAPIName
//	req.FieldAPIName = fieldAPIName
//
//	ctx, cancel, namespace, err := r.pre(ctx, appCtx, cConstants.GetField)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//	req.Namespace = namespace
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.GetObjectField(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	if err := cUtils.JsonUnmarshalBytes([]byte(resp.GetData()), &result); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func (r *RequestRpc) MGetUserSettings(ctx context.Context, appCtx *structs.AppCtx, userIDList []int64) (result []*structs.UserSetting, err error) {
//	req := identity.NewBatchGetUserSettingRequest()
//	req.UserIDs = userIDList
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.MGetUserSettings)
//	if err != nil {
//		return nil, err
//	}
//	defer cancel()
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := cli.BatchGetUserSetting(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return nil, err
//	}
//
//	if err := cUtils.JsonUnmarshalBytes([]byte(resp.GetUserSettings()), &result); err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func (r *RequestRpc) InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error {
//	sysParams, bizParams, err := reqCommon.BuildInvokeParamAndContext(ctx, params)
//	if err != nil {
//		return err
//	}
//
//	req := cloudfunction.NewInvokeFunctionWithAuthRequest()
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.InvokeFuncWithAuth)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	ctx = utils.SetUserMetaInfoToContext(ctx, appCtx)
//
//	namespace, err := utils.GetNamespace(ctx, appCtx)
//	if err != nil {
//		return err
//	}
//	req.Namespace = namespace
//	req.ApiName = cUtils.StringPtr(apiName)
//	req.Context = cUtils.StringPtr(sysParams)
//	req.Params = cUtils.StringPtr(bizParams)
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.InvokeFunctionWithAuth(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	var logid string
//	if req.Base != nil {
//		logid = req.Base.LogID
//	}
//
//	if resp.Result_ != nil {
//		data := []byte(*resp.Result_)
//		code := gjson.GetBytes(data, "code").String()
//		if code != "0" {
//			msg := gjson.GetBytes(data, "msg").String()
//			return cExceptions.InvalidParamError("%v ([%v] %v)", msg, code, logid)
//		}
//
//		dataRaw := gjson.GetBytes(data, "data").Raw
//		if len(dataRaw) > 0 {
//			if err := cUtils.JsonUnmarshalBytes([]byte(dataRaw), result); err != nil {
//				return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
//			}
//		}
//		return nil
//	}
//
//	return nil
//}
//
//func (r *RequestRpc) InvokeFunctionAsync(ctx context.Context, appCtx *structs.AppCtx, apiName string, params map[string]interface{}) (int64, error) {
//	sysParams, bizParams, err := reqCommon.BuildInvokeParamAndContext(ctx, params)
//	if err != nil {
//		return 0, err
//	}
//
//	req := cloudfunction.NewCreateAsyncTaskRequest()
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.CreateAsyncTask)
//	if err != nil {
//		return 0, err
//	}
//	defer cancel()
//
//	ctx = utils.SetUserMetaInfoToContext(ctx, appCtx)
//	namespace, err := utils.GetNamespace(ctx, appCtx)
//	if err != nil {
//		return 0, err
//	}
//	req.Namespace = namespace
//	req.APIAlias = apiName
//	req.Context = sysParams
//	req.TriggerType = "workflow"
//	req.Params = cUtils.StringPtr(bizParams)
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return 0, err
//	}
//
//	resp, err := cli.CreateAsyncTask(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return 0, err
//	}
//
//	return resp.TaskID, nil
//}
