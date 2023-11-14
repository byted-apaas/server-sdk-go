// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package openapi

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	cConstants "github.com/byted-apaas/server-common-go/constants"
	cHttp "github.com/byted-apaas/server-common-go/http"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/version"
)

var (
	openapiClientOnce sync.Once
	openapiClient     *cHttp.HttpClient
)

func getOpenapiClient() *cHttp.HttpClient {
	openapiClientOnce.Do(func() {
		openapiClient = &cHttp.HttpClient{
			Type: cHttp.OpenAPIClient,
			Client: http.Client{
				Transport: &http.Transport{
					DialContext:         cHttp.TimeoutDialer(cConstants.HttpClientDialTimeoutDefault, 0),
					TLSHandshakeTimeout: cConstants.HttpClientTLSTimeoutDefault,
					MaxIdleConns:        1000,
					MaxIdleConnsPerHost: 10,
					IdleConnTimeout:     60 * time.Second,
				},
			},
			MeshClient: &http.Client{
				Transport: &http.Transport{
					DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
						unixAddr, err := net.ResolveUnixAddr("unix", cUtils.GetSocketAddr())
						if err != nil {
							return nil, err
						}
						return net.DialUnix("unix", nil, unixAddr)
					},
					TLSHandshakeTimeout: cConstants.HttpClientTLSTimeoutDefault,
					MaxIdleConns:        1000,
					MaxIdleConnsPerHost: 10,
					IdleConnTimeout:     60 * time.Second,
				},
			},
			FromSDK: version.GetServerSDKInfo(),
		}
	})
	return openapiClient
}
