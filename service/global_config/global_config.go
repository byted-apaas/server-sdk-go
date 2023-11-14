// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package global_config

import (
	"context"
	"time"

	"github.com/muesli/cache2go"

	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	"github.com/byted-apaas/server-sdk-go/common/constants"
	"github.com/byted-apaas/server-sdk-go/common/structs"
	"github.com/byted-apaas/server-sdk-go/request"
)

func GetVar(ctx context.Context, key string) (string, error) {
	return GetVariable(ctx, nil, key)
}

func GetVariable(ctx context.Context, appCtx *structs.AppCtx, key string) (string, error) {
	// from local cache
	valuePtr := getVariableFromLocalCache(key)
	if valuePtr != nil {
		return *valuePtr, nil
	}

	// from remote
	keyToValue, err := request.GetInstance(ctx).GetAllGlobalConfig(ctx, appCtx)
	if err != nil {
		return "", err
	}

	// save cache
	addVariablesToLocalCache(keyToValue)

	if value, ok := keyToValue[key]; ok {
		return value, nil
	}
	return "", cExceptions.InvalidParamError("The global config (%s) does not exist", key)
}

func getVariableFromLocalCache(key string) *string {
	cacheTable := cache2go.Cache(constants.GlobalVariableCacheTableKey)
	cacheItem, err := cacheTable.Value(key)
	if err != nil {
		return nil
	}
	if value, ok := cacheItem.Data().(string); ok {
		return &value
	}
	return nil
}

func addVariablesToLocalCache(keyToValue map[string]string) {
	cacheTable := cache2go.Cache(constants.GlobalVariableCacheTableKey)

	for key, value := range keyToValue {
		cacheTable.Add(key, time.Minute, value)
	}
}
