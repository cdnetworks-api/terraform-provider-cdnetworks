package cdnetworks

import (
	"context"
	sdkCommon "github.com/cdnetworks-api/cdnetworks-sdk-go/common"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/connectivity"
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/cdn/domain"
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/monitor/rule"
	"github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/ssl/certificate"
	waapCustomizerule "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/waap/customizerule"
	waapDomain "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/waap/domain"
	waapRatelimit "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/waap/ratelimit"
	waapWhitelist "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/services/waap/whitelist"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	PROVIDER_SECRET_ID  = "CDNETWORKS_SECRET_ID"
	PROVIDER_SECRET_KEY = "CDNETWORKS_SECRET_KEY"
	PROVIDER_PROTOCOL   = "CDNETWORKS_PROTOCOL"
	PROVIDER_DOMAIN     = "CDNETWORKS_DOMAIN"
)

type CdnetworksClient struct {
	apiV3Conn *connectivity.CdnetworksClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the cdnetworks access key. It must be provided, but it can also be sourced from the `CDNETWORKS_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the cdnetworks secret key. It must be provided, but it can also be sourced from the `CDNETWORKS_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "https"),
				ValidateFunc: cdnetworksCommon.ValidateAllowedStringValue([]string{"http", "https"}),
				Description:  "(Optional)The protocol of the API request. Valid values: `http` and `https`. Default is `https`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "(Optional)The root domain of the API request.Default is `open.chinanetcenter.com`. It is optional",
			},
			"service_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(Optional)Security service type. Please enter a specific service type, if you purchase multiple security services.",
			},
			"contract_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(Optional)The id of contract, such as 40015677.Please enter a specific contract id.",
			},
			"item_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(Optional)The id of product, such as 10.Please enter a specific item id.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cdnetworks_cdn_domain":            domain.ResourceCdnDomain(),
			"cdnetworks_ssl_certificate":       certificate.ResourceSslCertificate(),
			"cdnetworks_waap_whitelist":        waapWhitelist.ResourceWaapWhitelist(),
			"cdnetworks_waap_customizerule":    waapCustomizerule.ResourceWaapCustomizeRule(),
			"cdnetworks_waap_ratelimit":        waapRatelimit.ResourceWaapRateLimit(),
			"cdnetworks_waap_domain_copy":      waapDomain.ResourceWaapDomainCopy(),
			"cdnetworks_waap_domain":           waapDomain.ResourceWaapDomain(),
			"cdnetworks_monitor_realtime_rule": rule.ResourceMonitorRealtimeRule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cdnetworks_cdn_domains":             domain.DataSourceCdnetworksCdnDomains(),
			"cdnetworks_cdn_domain_detail":       domain.DataSourceCdnetworksCdnDomainDetail(),
			"cdnetworks_ssl_certificate_detail":  certificate.DataSourceSslCertificateDetail(),
			"cdnetworks_ssl_certificates":        certificate.DataSourceSslCertificates(),
			"cdnetworks_waap_whitelists":         waapWhitelist.DataSourceWaapWhitelists(),
			"cdnetworks_waap_customizerules":     waapCustomizerule.DataSourceCustomizeRules(),
			"cdnetworks_waap_ratelimits":         waapRatelimit.DataSourceRateLimits(),
			"cdnetworks_waap_domains":            waapDomain.DataSourceWaapDomains(),
			"cdnw_monitor_realtime_rules_detail": rule.DataSourceMonitorRealtimeRuleDetail(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var (
		secretId    string
		secretKey   string
		protocol    string
		domain      string
		serviceType string
	)
	if v, ok := d.GetOk("secret_id"); ok {
		secretId = v.(string)
	}

	if v, ok := d.GetOk("secret_key"); ok {
		secretKey = v.(string)
	}
	if v, ok := d.GetOk("protocol"); ok {
		protocol = v.(string)
	}

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
	}

	if v, ok := d.GetOk("service_type"); ok {
		serviceType = v.(string)
	}

	var cdnetworksClient CdnetworksClient
	cdnetworksClient.apiV3Conn = &connectivity.CdnetworksClient{
		Credential:  sdkCommon.NewCredential(secretId, secretKey),
		HttpProfile: sdkCommon.NewHttpProfile(domain, protocol, serviceType),
	}

	return &cdnetworksClient, nil
}

// GetAPIV3Conn 返回访问云 API 的客户端连接对象
func (client *CdnetworksClient) GetAPIV3Conn() *connectivity.CdnetworksClient {
	return client.apiV3Conn
}
