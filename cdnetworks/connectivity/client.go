package connectivity

import (
	cdn "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cdn/domain"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
)

type CdnetworksClient struct {
	Credential  *common.Credential
	HttpProfile *common.HttpProfile

	cdnConn *cdn.Client
}

func (me *CdnetworksClient) UseCdnClient() *cdn.Client {
	if me.cdnConn != nil {
		return me.cdnConn
	}

	me.cdnConn, _ = cdn.NewClient(me.Credential, me.HttpProfile)

	return me.cdnConn
}
