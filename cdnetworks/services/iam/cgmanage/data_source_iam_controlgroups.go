package cgmanage

import (
	cgmanage "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"log"
	"time"
)

func DataSourceIamControlGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIamControlGroupsRead,
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"control_group_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Control Group Name",
						},
						"control_group_code": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Control Group Code",
						},
						"control_group_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Control Group ID",
						},
						"control_group_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Control Group Type",
						},
					},
				},
			},
		},
	}
}
func dataSourceIamControlGroupsRead(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("data_source.cdnetworks_iam_controlgroups.read")
	var diags diag.Diagnostics
	var response *cgmanage.QueryControlGroupListResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, 2*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().QueryControlGroupList()
		if err != nil {
			return resource.RetryableError(err)
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}

	codes := make([]string, 0, len(response.Data))
	dataList := make([]interface{}, len(response.Data))
	for i, d := range response.Data {
		if d != nil {
			dataList[i] = map[string]interface{}{
				"control_group_name": d.CONTROLGROUPNAME,
				"control_group_code": d.CONTROLGROUPCODE,
				"control_group_id":   d.CONTROLGROUPID,
				"control_group_type": d.CONTROLGROUPTYPE,
			}
			codes = append(codes, *d.CONTROLGROUPCODE)
		}
	}
	_ = data.Set("data", dataList)
	data.SetId(cdnetworksCommon.DataResourceIdsHash(codes))
	log.Printf("resource.cdnetworks_iam_controlgroups.read finish, requestId: %s", requestId)
	return diags
}
