package cgmanage

import (
	"log"
	"time"

	cgmanage "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/cgmanage"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

func ResourceIamControlGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIamControlGroupCreate,
		ReadContext:   resourceIamControlGroupRead,
		UpdateContext: resourceIamControlGroupUpdate,
		DeleteContext: resourceIamControlGroupDelete,
		Schema: map[string]*schema.Schema{
			"controlgroup_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Control Group Name",
			},
			"account_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Account object array, Used to specify accounts with permission.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"domain_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Domain array, Used to specify the domain contained in the Control Group",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceIamControlGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_controlgroup.create")
	var diags diag.Diagnostics
	request := &cgmanage.CreateControlGroupRequest{}

	// Set required fields
	if controlGroupName, ok := data.Get("controlgroup_name").(string); ok && controlGroupName != "" {
		request.ControlGroupName = &controlGroupName
	}

	// Process domain list
	if v, ok := data.Get("domain_list").([]interface{}); ok && len(v) > 0 {
		domains := make([]*string, 0, len(v))
		for _, domain := range v {
			d := domain.(string)
			domains = append(domains, &d)
		}
		request.DomainList = domains
	}

	// Process account list
	if v, ok := data.Get("account_list").([]interface{}); ok && len(v) > 0 {
		accounts := make([]*cgmanage.CreateControlGroupRequestAccountList, 0, len(v))
		for _, acc := range v {
			if loginName, ok := acc.(string); ok && loginName != "" {
				account := &cgmanage.CreateControlGroupRequestAccountList{}
				account.LoginName = &loginName
				accounts = append(accounts, account)
			}
		}
		request.AccountList = accounts
	}

	// Call API to create control group
	var response *cgmanage.CreateControlGroupResponse
	var requestId string
	var err error

	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().CreateControlGroup(request)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response == nil || response.Data == nil {
		data.SetId("")
		return nil
	}

	data.SetId(*response.Data.ControlGroupCode)

	log.Printf("resource.cdnetworks_iam_controlgroup.create finish, requestId: %s", requestId)
	return diags
}

func resourceIamControlGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_controlgroup.read")
	var diags diag.Diagnostics

	// Prepare request to query control group details
	request := &cgmanage.QueryCustomizedControlGroupByNameRequest{}
	var controlGroupCode = data.Id()
	request.ControlGroupCode = &controlGroupCode
	// Call API to get control group details
	var response *cgmanage.QueryCustomizedControlGroupByNameResponse
	var requestId string
	var err error

	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().QueryCustomizedControlGroupByName(request)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response == nil || response.Data == nil {
		data.SetId("")
		return nil
	}

	responseData := response.Data

	_ = data.Set("controlgroup_name", *responseData.ControlGroupName)

	// Set domain list
	if responseData.DomainList != nil && len(responseData.DomainList) > 0 {
		domains := make([]string, len(responseData.DomainList))
		for i, domain := range responseData.DomainList {
			domains[i] = *domain
		}
		_ = data.Set("domain_list", domains)
	}

	// Set account list
	if responseData.AccountNameList != nil && len(responseData.AccountNameList) > 0 {
		accounts := make([]string, len(responseData.AccountNameList))
		for i, account := range responseData.AccountNameList {
			accounts[i] = *account
		}
		_ = data.Set("account_list", accounts)
	}

	log.Printf("resource.cdnetworks_iam_controlgroup.read finish, requestId: %s", requestId)
	return diags
}

func resourceIamControlGroupUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_controlgroup.update")
	var diags diag.Diagnostics
	request := &cgmanage.EditControlGroupByCoverRequest{}

	controlGroupCode := data.Id()
	request.ControlGroupCode = &controlGroupCode

	if controlGroupName, ok := data.Get("controlgroup_name").(string); ok && controlGroupName != "" {
		request.ControlGroupName = &controlGroupName
	}

	// Process domain list
	if v, ok := data.Get("domain_list").([]interface{}); ok && len(v) > 0 {
		domains := make([]*string, 0, len(v))
		for _, domain := range v {
			d := domain.(string)
			domains = append(domains, &d)
		}
		request.DomainList = domains
	} else {
		// 不需要支持传null当作不修改，传null或者[]均当作清空，需要传[]给W+接口
		request.DomainList = []*string{}
	}

	// Process account list
	if v, ok := data.Get("account_list").([]interface{}); ok && len(v) > 0 {
		accounts := make([]*cgmanage.EditControlGroupByCoverRequestAccountList, 0, len(v))
		for _, acc := range v {
			if loginName, ok := acc.(string); ok && loginName != "" {
				account := &cgmanage.EditControlGroupByCoverRequestAccountList{}
				account.LoginName = &loginName
				accounts = append(accounts, account)
			}
		}
		request.AccountList = accounts
	} else {
		// 不需要支持传null当作不修改，传null或者[]均当作清空，需要传[]给W+接口
		request.AccountList = []*cgmanage.EditControlGroupByCoverRequestAccountList{}
	}

	// Call API to update control group
	var response *cgmanage.EditControlGroupByCoverResponse
	var requestId string
	var err error

	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().EditControlGroupByCover(request)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	if response == nil || response.Data == nil {
		data.SetId("")
		return nil
	}

	log.Printf("resource.cdnetworks_iam_controlgroup.update finish, requestId: %s", requestId)
	return diags
}

func resourceIamControlGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_controlgroup.delete")
	var diags diag.Diagnostics

	// Prepare request to delete control group
	paths := &cgmanage.DeleteControlGroupPaths{}
	controlGroupCode := data.Id()
	paths.ControlGroupCode = &controlGroupCode

	// Call API to delete control group
	var response *cgmanage.DeleteControlGroupResponse
	var requestId string
	var err error

	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseIamCgManageClient().DeleteControlGroup(paths)
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
		data.SetId("")
		return nil
	}

	log.Printf("resource.cdnetworks_iam_controlgroup.delete finish, requestId: %s", requestId)
	return diags
}
