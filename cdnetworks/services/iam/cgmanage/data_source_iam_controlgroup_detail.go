package cgmanage

import (
	"context"
	cgmanage "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"time"
)

func DataSourceIamControlGroupDetail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamControlGroupDetailRead,
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controlgroup_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Control Group Name",
						},
						"account_list": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Account object array, Used to specify accounts with permission.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"domain_list": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Domain array, Used to specify the domain contained in the Control Group",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"controlgroup_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Control Group Name",
			},
		},
	}
}
func dataSourceIamControlGroupDetailRead(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("data_source.cdnetworks_iam_controlgroup_detail.read")
	var diags diag.Diagnostics
	var controlGroupName string
	if v, ok := data.Get("controlgroup_name").(string); ok {
		controlGroupName = v
	}
	request := &cgmanage.QueryCustomizedControlGroupByNameRequest{}
	request.ControlGroupName = &controlGroupName

	var response *cgmanage.QueryCustomizedControlGroupByNameResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().QueryCustomizedControlGroupByName(request)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	dataMap := make(map[string]interface{})
	if response.Data != nil {
		domainList := make([]interface{}, 0)
		if response.Data.DomainList != nil {
			for _, domain := range response.Data.DomainList {
				domainList = append(domainList, *domain)
			}
		}
		dataMap["domain_list"] = domainList

		accountNameList := make([]interface{}, 0)
		if response.Data.AccountNameList != nil {
			for _, account := range response.Data.AccountNameList {
				accountNameList = append(accountNameList, *account)
			}
		}
		dataMap["account_list"] = accountNameList

		dataMap["controlgroup_name"] = response.Data.ControlGroupName
	}
	_ = data.Set("data", []interface{}{dataMap})
	data.SetId(*response.Data.ControlGroupCode)
	log.Printf("resource.cdnetworks_iam_controlgroup_detail.read finish, requestId: %s", requestId)
	return diags
}
