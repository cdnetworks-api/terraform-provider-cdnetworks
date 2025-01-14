package common

import (
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/connectivity"
)

// ProviderMeta Provider 元信息
type ProviderMeta interface {
	// GetAPIV3Conn 返回访问云 API 的客户端连接对象
	GetAPIV3Conn() *connectivity.CdnetworksClient
}
