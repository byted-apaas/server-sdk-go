// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package structs

import (
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/service/data/field_type/faassdk"
	"context"
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
	Criterion                interface{}           `json:"criterion"`
	Order                    []*Order              `json:"order"`
	FieldApiNames            []string              `json:"field_api_names"`
	Offset                   int64                 `json:"offset"`
	Limit                    int64                 `json:"limit"`
	NeedTotalCount           bool                  `json:"need_total_count"`
	IgnoreBackLookupField    bool                  `json:"ignore_back_lookup_field"`
	NeedFilterUserPermission bool                  `json:"need_filter_user_permission"`
	FuzzySearch              *FuzzySearch          `json:"fuzzySearch"`
	ProcessAuthFieldType     *ProcessAuthFieldType `json:"process_auth_field_type"`
}

type ProcessAuthFieldType int64

const (
	ProcessAuthFieldType_Default ProcessAuthFieldType = iota
	ProcessAuthFieldType_BothResult
	ProcessAuthFieldType_SliceResult
	ProcessAuthFieldType_MapResult
)

type GetRecordsReqParamV2 struct {
	Limit       int64         `json:"limit"`
	Offset      int64         `json:"offset"`
	Fields      []string      `json:"fields"`
	QuickSearch string        `json:"quickSearch"`
	Filter      []interface{} `json:"filter"`
	Sort        []*Order      `json:"sort"`
	Count       *bool         `json:"count"`
}

type GetRecordsReqParamV3 struct {
	PageSize             int64                 `json:"page_size"` // 是期望服务端返回的条目个数，不填则取默认值，默认值为 500， 对应 v2 的 limit
	Offset               int64                 `json:"offset"`    // 偏移量，从 0 开始
	Select               []string              `json:"select"`    // 筛选的字段列表
	Filter               interface{}           `json:"filter"`    // Criterion
	OrderBy              []*Order              `json:"order_by"`
	ProcessAuthFieldType *ProcessAuthFieldType `json:"process_auth_field_type"`
	GroupBy              []GroupByItem         `json:"group_by"`
	NeedTotalCount       bool                  `json:"need_total_count"` // 是否需要返回符合条件的记录总数
	DataVersion          string                `json:"data_version"`     // v2 or v3
}

type GetSingleRecordReqParamV3 struct {
	Select      []string `json:"select"`       // 筛选的字段列表
	DataVersion string   `json:"data_version"` // v2 or v3
}

type GroupByItem struct {
	Field string `json:"field"`
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

type FindStreamData struct {
	Records      interface{} `json:"records"`
	UnauthFields [][]string  `json:"unauthFields"`
}

type FindStreamParam struct {
	IDGetter  func(record interface{}) (id int64, err error)
	Handler   func(ctx context.Context, data *FindStreamData) (err error)
	PageLimit int64
}
