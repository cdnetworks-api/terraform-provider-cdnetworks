package user

import (
	"log"
	"time"

	"github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/usermanage"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

func ResourceUserInfo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"login_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User login name",
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User display name",
			},
			"status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "User status. Options: 1-active, 0-inactive",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User email address",
			},
			"console_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether console access is enabled. Options: 1-enabled, 0-disabled",
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_user.create")
	var diags diag.Diagnostics
	request := &usermanage.AddSubAccountRequest{}

	if loginName, ok := data.Get("login_name").(string); ok && loginName != "" {
		request.LoginName = &loginName
	} else {
		log.Printf("login_name is required")
		return diags
	}
	if displayName, ok := data.Get("display_name").(string); ok && displayName != "" {
		request.DisplayName = &displayName
	} else {
		log.Printf("display_name is required")
		return diags
	}
	if status, ok := data.Get("status").(int); ok {
		request.Status = &status
	}
	if email, ok := data.Get("email").(string); ok && email != "" {
		request.Email = &email
	} else {
		log.Printf("email is required")
		return diags
	}
	if consoleEnable, ok := data.Get("console_enable").(int); ok && (consoleEnable == 0 || consoleEnable == 1) {
		request.ConsoleEnable = &consoleEnable
	}

	var response *usermanage.AddSubAccountResponse
	var err error
	var requestId string
	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseUserManageClient().CreateUser(request)
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
	} else {
		data.SetId(*request.LoginName)
	}
	_ = data.Set("login_name", *request.LoginName)
	log.Printf("resource.cdnetworks_iam_user.create success")
	log.Printf("requestId: %s", requestId)
	time.Sleep(2 * time.Second)
	return resourceUserRead(ctx, data, meta)
}

func resourceUserRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_user.read")
	var diags diag.Diagnostics
	request := &usermanage.QuerySubAccountInfoRequest{}

	var response *usermanage.QuerySubAccountInfoResponse
	if loginName, ok := data.Get("login_name").(string); ok && loginName != "" {
		path := &usermanage.QuerySubAccountInfoPaths{
			LoginName: &loginName,
		}
		var err error
		var requestId string
		err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
			requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseUserManageClient().QueryUser(request, path)
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
		data.SetId(*response.Data.LoginName)
		_ = data.Set("login_name", response.Data.LoginName)
		_ = data.Set("display_name", response.Data.DisplayName)
		_ = data.Set("status", response.Data.Status)
		_ = data.Set("email", response.Data.Email)
		_ = data.Set("console_enable", response.Data.ConsoleEnable)
		log.Printf("resource.cdnetworks_iam_user.read success, requestId: %s", requestId)
		return diags
	} else {
		log.Printf("login_name is required")
		return diags
	}

}

func resourceUserUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_user.update")
	var diags diag.Diagnostics
	request := &usermanage.UpdateSubAccountRequest{}
	if loginName, ok := data.Get("login_name").(string); ok && loginName != "" {
		request.LoginName = &loginName
	}
	if displayName, ok := data.Get("display_name").(string); ok && displayName != "" {
		request.DisplayName = &displayName
	}
	if status, ok := data.Get("status").(int); ok {
		request.Status = &status
	}
	if email, ok := data.Get("email").(string); ok {
		request.Email = &email
	}
	if consoleEnable, ok := data.Get("console_enable").(int); ok {
		request.ConsoleEnable = &consoleEnable
	}

	var response *usermanage.UpdateSubAccountResponse
	var err error
	var requestId string
	err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseUserManageClient().EditUser(request)
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
	data.SetId(*request.LoginName)
	log.Printf("resource.cdnetworks_iam_user.update success")
	log.Printf("requestId: %s", requestId)
	time.Sleep(2 * time.Second)
	return resourceUserRead(ctx, data, meta)
}

func resourceUserDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_iam_user.delete")
	var diags diag.Diagnostics
	request := &usermanage.DeleteSubAccountRequest{}

	var response *usermanage.DeleteSubAccountResponse
	if loginName, ok := data.Get("login_name").(string); ok && loginName != "" {
		path := &usermanage.DeleteSubAccountPaths{
			LoginName: &loginName,
		}
		var err error
		var requestId string
		err = resource.RetryContext(ctx, time.Duration(2)*time.Minute, func() *resource.RetryError {
			requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseUserManageClient().DeleteUser(request, path)
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
		log.Printf("resource.cdnetworks_iam_policy.delete success")
		log.Printf("requestId: %s", requestId)
		return diags
	} else {
		log.Printf("login_name is required")
		return diags
	}

}
