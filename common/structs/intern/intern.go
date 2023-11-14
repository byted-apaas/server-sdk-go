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

type TenantAccessToken struct {
	Expire            int64  `json:"expire"`
	TenantAccessToken string `json:"tenantAccessToken"`
	AppID             string `json:"appId"`
}

type AppAccessToken struct {
	Expire         int64  `json:"expire"`
	AppAccessToken string `json:"appAccessToken"`
	AppID          string `json:"appId"`
}

type UnauthPermissionInfo struct {
	UnauthFieldMap   map[int64][]string `json:"unauth_field_map"`
	UnauthFieldSlice [][]string         `json:"unauth_field_slice"`
}

type SetSystemMod int64

const (
	SetSystemMod_NotSet SetSystemMod = iota
	SetSystemMod_FullSet
	SetSystemMod_Other
)

type CommitSetSystemMod int64

const (
	CommitSetSystemMod_NotSet      CommitSetSystemMod = 0
	CommitSetSystemMod_SysFieldSet CommitSetSystemMod = 1
)
