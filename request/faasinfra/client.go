// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package faasinfra

import (
	"net/http"
	"sync"
	"time"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cHttp "github.com/byted-apaas/server-common-go/http"
)

var (
	fsInfraClientOnce sync.Once
	fsInfraClient     *cHttp.HttpClient
)

func getFaaSInfraClient() *cHttp.HttpClient {
	fsInfraClientOnce.Do(func() {
		fsInfraClient = &cHttp.HttpClient{
			Type: cHttp.FaaSInfraClient,
			Client: http.Client{
				Transport: &http.Transport{
					DialContext:         cHttp.TimeoutDialer(cConstants.HttpClientDialTimeoutDefault, 0),
					TLSHandshakeTimeout: cConstants.HttpClientTLSTimeoutDefault,
					MaxIdleConns:        1000,
					MaxIdleConnsPerHost: 10,
					IdleConnTimeout:     60 * time.Second,
				},
			},
		}
	})
	return fsInfraClient
}
