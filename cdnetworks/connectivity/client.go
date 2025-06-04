package connectivity

import (
	cdn "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cdn/domain"
	iamCgManage "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/common"
	monitorRule "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/monitor/rule"
	policy "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/policy"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ssl/certificate"
	userManage "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage"
	waapCustomizerule "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/customizerule"
	waapDomain "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/domain"
	waapRatelimit "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/ratelimit"
	waapShareCustomizerule "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/share-customizerule"
	waapShareWhitelist "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/share-whitelist"
	waapWhitelist "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/whitelist"
)

type CdnetworksClient struct {
	Credential  *common.Credential
	HttpProfile *common.HttpProfile

	cdnConn                    *cdn.Client
	sslCertificateConn         *certificate.Client
	waapWhitelistConn          *waapWhitelist.Client
	waapCustomizeruleConn      *waapCustomizerule.Client
	waapRatelimitConn          *waapRatelimit.Client
	waapDomainConn             *waapDomain.Client
	waapShareWhitelistConn     *waapShareWhitelist.Client
	waapShareCustomizeruleConn *waapShareCustomizerule.Client
	monitorRuleConn            *monitorRule.Client
	iamCgManageConn            *iamCgManage.Client
	policyConn                 *policy.Client
	userManageConn             *userManage.Client
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

func (me *CdnetworksClient) UseWaapShareWhitelistClient() *waapShareWhitelist.Client {
	if me.waapShareWhitelistConn != nil {
		return me.waapShareWhitelistConn
	}

	me.waapShareWhitelistConn, _ = waapShareWhitelist.NewClient(me.Credential, me.HttpProfile)

	return me.waapShareWhitelistConn
}

func (me *CdnetworksClient) UseWaapShareCustomizeruleClient() *waapShareCustomizerule.Client {
	if me.waapShareCustomizeruleConn != nil {
		return me.waapShareCustomizeruleConn
	}

	me.waapShareCustomizeruleConn, _ = waapShareCustomizerule.NewClient(me.Credential, me.HttpProfile)

	return me.waapShareCustomizeruleConn
}

func (me *CdnetworksClient) UseMonitorRuleClient() *monitorRule.Client {
	if me.monitorRuleConn != nil {
		return me.monitorRuleConn
	}

	me.monitorRuleConn, _ = monitorRule.NewClient(me.Credential, me.HttpProfile)

	return me.monitorRuleConn
}
func (me *CdnetworksClient) UseIamCgManageClient() *iamCgManage.Client {
	if me.iamCgManageConn != nil {
		return me.iamCgManageConn
	}

	me.iamCgManageConn, _ = iamCgManage.NewClient(me.Credential, me.HttpProfile)

	return me.iamCgManageConn
}

func (me *CdnetworksClient) UsePolicyClient() *policy.Client {
	if me.policyConn != nil {
		return me.policyConn
	}

	me.policyConn, _ = policy.NewClient(me.Credential, me.HttpProfile)

	return me.policyConn
}

func (me *CdnetworksClient) UsePolicyAttachmentClient() *userManage.Client {
	if me.userManageConn != nil {
		return me.userManageConn
	}

	me.userManageConn, _ = userManage.NewClient(me.Credential, me.HttpProfile)

	return me.userManageConn
}

func (me *CdnetworksClient) UseUserManageClient() *userManage.Client {
	if me.userManageConn != nil {
		return me.userManageConn
	}

	me.userManageConn, _ = userManage.NewClient(me.Credential, me.HttpProfile)

	return me.userManageConn
}
