// Package version defines version of server-sdk-go.
package version

import (
	"sync"

	cVersion "github.com/byted-apaas/server-common-go/version"
)

const Version = "v0.0.35-beta.12"

const SDKName = "byted-apaas/server-sdk-go"

type ServerSDKInfo struct{}

func (c *ServerSDKInfo) GetVersion() string {
	return Version
}

func (c *ServerSDKInfo) GetSDKName() string {
	return SDKName
}

var (
	serverSDKInfoOnce sync.Once
	serverSDKInfo     cVersion.ISDKInfo
)

func GetServerSDKInfo() cVersion.ISDKInfo {
	if serverSDKInfo == nil {
		serverSDKInfoOnce.Do(func() {
			serverSDKInfo = &ServerSDKInfo{}
		})
	}
	return serverSDKInfo
}
