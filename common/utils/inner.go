package utils

import (
	"context"
	"encoding/json"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cHttp "github.com/byted-apaas/server-common-go/http"
	log "github.com/byted-apaas/server-common-go/logger"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/common/structs/intern"
)

func GoWithRecover(ctx context.Context, f func() error) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				log.GetLogger(ctx).Errorf("panic: err is %+v, stack is %+v\n", e, string(buf))
			}
		}()
		if e := f(); e != nil {
			log.GetLogger(ctx).Errorf("[GoWithRecover] Do func failed, err: %+v\n", e)
		}
	}()
}

func SetCtx(ctx context.Context, appCtx *structs.AppCtx, method string) context.Context {
	ctx = SetAppConfToCtx(ctx, appCtx)
	if ctx == nil {
		ctx = context.Background()
	}
	ctx = cUtils.SetApiTimeoutMethodToCtx(ctx, method)
	return ctx
}

func SetAppConfToCtx(ctx context.Context, appCtx *structs.AppCtx) context.Context {
	if appCtx == nil || appCtx.Mode != structs.AppModeOpenSDK {
		return ctx
	}

	if ctx == nil {
		ctx = context.Background()
	}

	targetEnv := appCtx.GetEnv()
	conf, _ := cConstants.EnvConfMap[targetEnv.String()]

	ctx = cHttp.SetCredentialToCtx(ctx, appCtx.Credential)
	ctx = context.WithValue(ctx, cConstants.CtxKeyInnerAPIPSM, conf.InnerAPIPSM)
	ctx = context.WithValue(ctx, cConstants.CtxKeyOpenapiDomain, conf.OpenAPIDomain)
	ctx = context.WithValue(ctx, cConstants.CtxKeyFaaSInfraDomain, conf.FaaSInfraDomain)
	ctx = context.WithValue(ctx, cConstants.CtxKeyAGWDomain, conf.InnerAPIDomain)
	if strings.HasSuffix(targetEnv.String(), "boe") {
		ctx = cUtils.SetEnvBoeToCtx(ctx, "boe")
	}

	return ctx
}

func GetNamespace(ctx context.Context, appCtx *structs.AppCtx) (string, error) {
	if appCtx != nil && appCtx.Credential != nil && appCtx.Mode == structs.AppModeOpenSDK {
		tenant, err := appCtx.Credential.GetTenantInfo(ctx)
		if err != nil {
			return "", cExceptions.InternalError("GetTenantInfo failed, err: %v", err)
		}
		return tenant.Namespace, nil
	}
	return cUtils.GetNamespace(), nil
}

func GetTenantName(ctx context.Context, appCtx *structs.AppCtx) (string, error) {
	if appCtx != nil && appCtx.Credential != nil && appCtx.Mode == structs.AppModeOpenSDK {
		tenant, err := appCtx.Credential.GetTenantInfo(ctx)
		if err != nil {
			return "", cExceptions.InternalError("GetTenantInfo failed, err: %v", err)
		}
		return tenant.Name, nil
	}
	return cUtils.GetTenantName(), nil
}

func ParseInt64(v interface{}) (int64, error) {
	switch v.(type) {
	case json.Number:
		s, err := v.(json.Number).Int64()
		if err != nil {
			return s, cExceptions.InvalidParamError("ParseInt64 failed, err: %+v", err)
		}
		return s, nil
	case int:
		return int64(v.(int)), nil
	case int32:
		return int64(v.(int32)), nil
	case int64:
		return v.(int64), nil
	case float64:
		return int64(v.(float64)), nil
	case float32:
		return int64(v.(float32)), nil
	default:
		return 0, cExceptions.InvalidParamError("Cannot convert (%v) to int64", v)
	}
}

func ParseMultilingualNewToOld(new structs.Multilingual) (old intern.Multilingual) {
	if new.Zh != "" {
		old = append(old, &intern.MultilingualItem{
			LanguageCode: intern.LanguageCodeZH,
			Text:         new.Zh,
		})
	}
	if new.En != "" {
		old = append(old, &intern.MultilingualItem{
			LanguageCode: intern.LanguageCodeEN,
			Text:         new.En,
		})
	}
	return old
}

func ParseMapToFlowVariable(params map[string]interface{}) intern.ExecuteFlowVariables {
	variables := intern.ExecuteFlowVariables{}
	for k, v := range params {
		variables = append(variables, intern.ExecuteFlowVariable{
			APIName: k,
			Value:   v,
		})
	}
	return variables
}

func ParseFlowVariableToMap(variables intern.ExecuteFlowVariables) map[string]interface{} {
	res := make(map[string]interface{})
	for _, variable := range variables {
		res[variable.APIName] = variable.Value
	}
	return res
}

func SetUserMetaInfoToContext(ctx context.Context, appCtx *structs.AppCtx) context.Context {
	if appCtx.IsOpenSDK() {
		return ctx
	}
	ctx = context.WithValue(ctx, cConstants.HttpHeaderKeyUser, strconv.FormatInt(cUtils.GetUserIDFromCtx(ctx), 10))
	return ctx
}

func StrInStrs(strs []string, str string) bool {
	for _, v := range strs {
		if str == v {
			return true
		}
	}
	return false
}

func ParseBatchResult(resp *structs.BatchResult, result interface{}) error {
	if resp == nil {
		return nil
	}

	switch result.(type) {
	case *structs.BatchResult:
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(*resp))
	case **structs.BatchResult:
		reflect.ValueOf(result).Elem().Elem().Set(reflect.ValueOf(*resp))
	default:
		return cExceptions.InvalidParamError("the type of result should be *structs.BatchResult or **structs.BatchResult, but %T", result)
	}
	return nil
}
