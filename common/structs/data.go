// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import (
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
)

type RecordID struct {
	ID int64 `json:"_id"`
}

type TransactionRecordID struct {
	ID interface{} `json:"_id"`
}

type RecordsResult struct {
	Records []interface{} `json:"records"`
	Total   int64         `json:"total"`
}

type FuzzySearch struct {
	Keyword       string   `json:"keyword"`
	FieldAPINames []string `json:"fieldAPINames"`
}

type GetRecordsReqParam struct {
	Criterion                interface{}                     `json:"criterion"`
	Order                    []*Order                        `json:"order"`
	FieldApiNames            []string                        `json:"field_api_names"`
	Offset                   int64                           `json:"offset"`
	Limit                    int64                           `json:"limit"`
	NeedTotalCount           bool                            `json:"need_total_count"`
	IgnoreBackLookupField    bool                            `json:"ignore_back_lookup_field"`
	NeedFilterUserPermission bool                            `json:"need_filter_user_permission"`
	FuzzySearch              *FuzzySearch                    `json:"fuzzySearch"`
	ProcessAuthFieldType     *constants.ProcessAuthFieldType `json:"process_auth_field_type"`
}

type GetRecordsReqParamV2 struct {
	Limit       int64         `json:"limit"`
	Offset      int64         `json:"offset"`
	Fields      []string      `json:"fields"`
	QuickSearch string        `json:"quickSearch"`
	Filter      []interface{} `json:"filter"`
	Sort        []*Order      `json:"sort"`
	Count       *bool         `json:"count"`
}

type BatchCreateRecord struct {
	RecordIDs []int64 `json:"record_ids"`
}

type BatchCreateRecordV2 struct {
	RecordIDs []int64 `json:"_ids"`
}

type AsyncTaskResult struct {
	TaskID int64 `json:"taskID"`
}

type Order struct {
	Field     string `json:"field"`     // 字段的 apiName
	Direction string `json:"direction"` // 排序方向，取值为：asc 或 desc
	Type      string `json:"type"`      // 传空串即可，不传底层报错
}

type OqlResultSet struct {
	Rows interface{} `json:"rows"`
}

type MessageBody struct {
	Icon        string                `json:"icon"` // "success" | "error" | "progress" | "info"
	Percent     int64                 `json:"percent"`
	TargetUsers []int64               `json:"targetUsers"`
	Title       *faassdk.Multilingual `json:"title"`
	Detail      *faassdk.Multilingual `json:"detail"`
}

type MessageParam struct {
	NotifyModelKey     string                 `json:"NotifyModelKey"`
	ReceiverIDs        []int64                `json:"ReceiverIDs"`
	ParamsRawData      interface{}            `json:"ParamsRawData"`
	Percent            interface{}            `json:"Percent"`
	ModelKeySearchData map[string]interface{} `json:"ModelKeySearchData"`
}

type TransactionOperation struct {
	OperationType constants.OperationType `json:"operationType"`
	ObjectAPIName string                  `json:"objectApiName"`
	Input         string                  `json:"input"`
}

type GlobalConfigResult struct {
	Configs []struct {
		BizType string `json:"biz_type"`
		Extra   struct {
			AccessSpecifier string `json:"access_specifier"`
		} `json:"extra"`
		Key       string `json:"key"`
		NameSpace string `json:"name_space"`
		Value     string `json:"value"`
	} `json:"configs"`
	Total int `json:"total"`
}

type BatchResult struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []BatchResultData `json:"data"`
}

func (b *BatchResult) HasError() bool {
	if b == nil {
		return false
	}

	for _, d := range b.Data {
		if !d.Success {
			return true
		}
	}
	return false
}

type BatchResultData struct {
	Success bool                   `json:"success"`
	ID      int64                  `json:"_id"`
	Errors  []BatchResultDataError `json:"errors"`
}

type BatchResultDataError struct {
	Code string `json:"code"`
}

type UnauthPermissionInfo struct {
	UnauthFieldMap   map[int64][]string `json:"unauth_field_map"`
	UnauthFieldSlice [][]string         `json:"unauth_field_slice"`
}
