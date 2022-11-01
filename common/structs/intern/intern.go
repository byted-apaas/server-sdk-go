package intern

const (
	LanguageCodeZH = 2052
	LanguageCodeEN = 1033
)

// Multilingual 多语结构
type Multilingual []*MultilingualItem

type MultilingualItem struct {
	LanguageCode int    `json:"language_code"`
	Text         string `json:"text"`
}

type ExecuteFlowVariable struct {
	APIName string      `json:"api_name"`
	Value   interface{} `json:"value"`
}
type ExecuteFlowVariables []ExecuteFlowVariable

type FlowExecuteResult struct {
	ExecutionID int64                `json:"executionId"`
	Status      string               `json:"status"`
	OutParams   ExecuteFlowVariables `json:"outParams"`
	ErrCode     *string              `json:"errCode"`
	ErrMsg      *string              `json:"errMsg"`
}

type ExecutionInfo struct {
	Status    string               `json:"status"`
	OutParams ExecuteFlowVariables `json:"outParams"`
	ErrCode   *string              `json:"errCode"`
	ErrMsg    *string              `json:"errMsg"`
}
