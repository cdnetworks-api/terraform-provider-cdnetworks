package connectivity

import (
	cdn "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cdn/domain"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ssl/certificate"
	waapCustomizerule "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/customizerule"
	waapDomain "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/domain"
	waapRatelimit "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/ratelimit"
	waapWhitelist "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/whitelist"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common"
)

type CdnetworksClient struct {
	Credential  *common.Credential
	HttpProfile *common.HttpProfile

	cdnConn            *cdn.Client
	sslCertificateConn *certificate.Client
	waapWhitelistConn     *waapWhitelist.Client
	waapCustomizeruleConn *waapCustomizerule.Client
	waapRatelimitConn     *waapRatelimit.Client
	waapDomainConn        *waapDomain.Client
}

func (me *CdnetworksClient) UseCdnClient() *cdn.Client {
	if me.cdnConn != nil {
		return me.cdnConn
	}

	me.cdnConn, _ = cdn.NewClient(me.Credential, me.HttpProfile)

	return me.cdnConn
}

func (me *CdnetworksClient) UseSslCertificateClient() *certificate.Client {
	if me.sslCertificateConn != nil {
		return me.sslCertificateConn
	}

	me.sslCertificateConn, _ = certificate.NewClient(me.Credential, me.HttpProfile)

	return me.sslCertificateConn
}

func (me *CdnetworksClient) UseWaapWhitelistClient() *waapWhitelist.Client {
	if me.waapWhitelistConn != nil {
		return me.waapWhitelistConn
	}

	me.waapWhitelistConn, _ = waapWhitelist.NewClient(me.Credential, me.HttpProfile)

	return me.waapWhitelistConn
}

func (me *CdnetworksClient) UseWaapCustomizeruleClient() *waapCustomizerule.Client {
	if me.waapCustomizeruleConn != nil {
		return me.waapCustomizeruleConn
	}

	me.waapCustomizeruleConn, _ = waapCustomizerule.NewClient(me.Credential, me.HttpProfile)

	return me.waapCustomizeruleConn
}

func (me *CdnetworksClient) UseWaapRatelimitClient() *waapRatelimit.Client {
	if me.waapRatelimitConn != nil {
		return me.waapRatelimitConn
	}

	me.waapRatelimitConn, _ = waapRatelimit.NewClient(me.Credential, me.HttpProfile)

	return me.waapRatelimitConn
}

func (me *CdnetworksClient) UseWaapDomainClient() *waapDomain.Client {
	if me.waapDomainConn != nil {
		return me.waapDomainConn
	}

	me.waapDomainConn, _ = waapDomain.NewClient(me.Credential, me.HttpProfile)

	return me.waapDomainConn
}
