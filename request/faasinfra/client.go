// Copyright 2022 ByteDance Ltd. and/or its affiliates
// SPDX-License-Identifier: MIT

package faasinfra

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
			FromSDK: version.GetServerSDKInfo(),
		}
	})
	if cUtils.EnableMesh() {
		fsInfraClient.MeshClient = &http.Client{
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
		}
	}
	return fsInfraClient
}
