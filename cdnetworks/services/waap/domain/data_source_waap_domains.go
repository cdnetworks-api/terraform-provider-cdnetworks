package domain

import (
	"context"
	"fmt"
	waapDomain "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/waap/domain"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

func DataSourceWaapDomains() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWaapDomainsRead,

		Schema: map[string]*schema.Schema{
			"defend_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protection status, If not specified, it means all the protection status.<br/>PROTECTING: Protecting<br/>UNPROTECTED: Unprotected",
			},
			"domain_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Hostname list, if not specified, it means all the hostnames of the account.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"dms_defend_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DDoS protection switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"rate_limit_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rate limiting switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"block_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP/Geo blocking switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"waf_defend_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "WAF protection switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"intelligence_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Threat intelligence switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"whitelist_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Whitelist switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"bot_manage_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Bot management switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"customize_rule_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom rules switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"api_defend_switch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "API security switch, if not specified, it means all the status.<br/>ON: Enabled<br/>OFF: Disabled",
			},
			"data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Created time, format: yyyy-MM-dd HH:mm:ss.",
						},
						"deploy_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Deployment status.<br/>DEPLOYING: Publishing<br/>SUCCESS: Success",
						},
						"block_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP/Geo blocking switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"defend_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protection status.<br/>PROTECTING: Protecting<br/>UNPROTECTED: Unprotected",
						},
						"dms_defend_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DDoS protection switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"bot_manage_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bot management switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"customize_rule_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Custom rules switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"api_defend_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "API security switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"rate_limit_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rate limiting switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"whitelist_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whitelist switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"intelligence_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Threat intelligence switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
						"waf_defend_switch": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "WAF protection switch.<br/>ON: Enabled<br/>OFF: Disabled",
						},
					},
				},
			},
		},
	}
}

func dataSourceWaapDomainsRead(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("data_source.cdnetworks_waap_domains.read")

	var response *waapDomain.ListDomainInfoResponse
	var err error
	var diags diag.Diagnostics
	request := &waapDomain.ListDomainInfoRequest{}
	if v, ok := data.GetOk("domain_list"); ok {
		domainsList := v.([]interface{})
		domainsListStr := make([]*string, len(domainsList))
		for i, v := range domainsList {
			str := v.(string)
			domainsListStr[i] = &str
		}
		request.SetDomainList(domainsListStr)
	}
	if v, ok := data.GetOk("defend_status"); ok {
		request.SetDefendStatus(v.(string))
	}
	if v, ok := data.GetOk("dms_defend_switch"); ok {
		request.SetDmsDefendSwitch(v.(string))
	}
	if v, ok := data.GetOk("rate_limit_switch"); ok {
		request.SetRateLimitSwitch(v.(string))
	}
	if v, ok := data.GetOk("block_switch"); ok {
		request.SetBlockSwitch(v.(string))
	}
	if v, ok := data.GetOk("waf_defend_switch"); ok {
		request.SetWafDefendSwitch(v.(string))
	}
	if v, ok := data.GetOk("intelligence_switch"); ok {
		request.SetIntelligenceSwitch(v.(string))
	}
	if v, ok := data.GetOk("whitelist_switch"); ok {
		request.SetWhitelistSwitch(v.(string))
	}
	if v, ok := data.GetOk("bot_manage_switch"); ok {
		request.SetBotManageSwitch(v.(string))
	}
	if v, ok := data.GetOk("customize_rule_switch"); ok {
		request.SetCustomizeRuleSwitch(v.(string))
	}
	if v, ok := data.GetOk("api_defend_switch"); ok {
		request.SetApiDefendSwitch(v.(string))
	}

	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		_, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseWaapDomainClient().GetDomainList(request)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response == nil {
		return nil
	}
	if len(response.Data) == 0 {
		return nil
	}
	dataList := make([]interface{}, len(response.Data))
	ids := make([]string, 0, len(response.Data))
	for i, v := range response.Data {
		dataList[i] = map[string]interface{}{
			"id":                    v.Id,
			"domain":                v.Domain,
			"create_time":           v.CreateTime,
			"deploy_status":         v.DeployStatus,
			"block_switch":          v.BlockSwitch,
			"defend_status":         v.DefendStatus,
			"dms_defend_switch":     v.DmsDefendSwitch,
			"bot_manage_switch":     v.BotManageSwitch,
			"customize_rule_switch": v.CustomizeRuleSwitch,
			"api_defend_switch":     v.ApiDefendSwitch,
			"rate_limit_switch":     v.RateLimitSwitch,
			"whitelist_switch":      v.WhitelistSwitch,
			"intelligence_switch":   v.IntelligenceSwitch,
			"waf_defend_switch":     v.WafDefendSwitch,
		}
		ids = append(ids, *v.Id)
	}
	if err := data.Set("data", dataList); err != nil {
		return diag.FromErr(fmt.Errorf("error setting data for resource: %s", err))
	}
	data.SetId(cdnetworksCommon.DataResourceIdsHash(ids))
	return diags
}
