package certificate

import (
	"context"
	certicate "github.com/cdnetworks-api/cdnetworks-sdk-go/cdnetworks/ssl/certificate"
	cdnetworksCommon "github.com/cdnetworks-api/terraform-provider-cdnetworks/cdnetworks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
	"time"
)

func ResourceSslCertificate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSslCertificateCreate,
		ReadContext:   resourceSslCertificateRead,
		UpdateContext: resourceSslCertificateUpdate,
		DeleteContext: resourceSslCertificateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate name",
			},
			"cert": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Certificate, PEM certificate, including CRT file and CA file.",
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private key of the certificate, PEM certificate.",
			},
			"certificate_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "certificate Id",
			},
		},
	}
}

func resourceSslCertificateCreate(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_ssl_certificate.create")

	var diags diag.Diagnostics
	request := &certicate.AddCertificateForTerraformRequest{}
	if name, ok := data.Get("name").(string); ok && name != "" {
		request.Name = &name
	}
	if certificate, ok := data.Get("cert").(string); ok && certificate != "" {
		request.Certificate = &certificate
	}
	if privateKey, ok := data.Get("key").(string); ok && privateKey != "" {
		request.PrivateKey = &privateKey
	}

	//start to create a domain in 2 minutes
	var response *certicate.AddCertificateForTerraformResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseSslCertificateClient().AddCertificate(request)
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
	var certificateIdStr = strconv.FormatInt(*response.Data.CertificateId, 10)
	data.Set("certificate_id", response.Data.CertificateId)
	data.SetId(certificateIdStr)
	log.Printf("resource.cdnetworks_ssl_certificate.create success")
	log.Printf("requestId: %s", requestId)
	time.Sleep(2 * time.Second)
	return resourceSslCertificateRead(context, data, meta)
}

func resourceSslCertificateRead(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_ssl_certificate.read")
	var diags diag.Diagnostics
	var certificateId, _ = strconv.ParseInt(data.Id(), 10, 64)
	var response *certicate.QueryCertificateForTerraformResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseSslCertificateClient().QueryCertificate(certificateId)
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
	data.Set("certificate_id", response.Data.CertificateId)
	data.Set("name", response.Data.Name)
	data.Set("cert", response.Data.Certificate)
	data.Set("key", response.Data.PrivateKey)
	log.Printf("resource.cdnetworks_ssl_certificate.read success")
	//打印requestId
	log.Printf("requestId: %s", requestId)
	return nil
}

func resourceSslCertificateUpdate(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_ssl_certificate.update")
	var diags diag.Diagnostics
	var certificateId, _ = strconv.ParseInt(data.Id(), 10, 64)
	request := &certicate.UpdateCertificateForTerraformRequest{}
	if name, ok := data.Get("name").(string); ok && name != "" {
		request.Name = &name
	}
	if certificate, ok := data.Get("cert").(string); ok && certificate != "" {
		request.Certificate = &certificate
	}
	if privateKey, ok := data.Get("key").(string); ok && privateKey != "" {
		request.PrivateKey = &privateKey
	}
	var response *certicate.UpdateCertificateForTerraformResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseSslCertificateClient().UpdateCertificate(certificateId, request)
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
	log.Printf("resource.cdnetworks_ssl_certificate.update success")
	log.Printf("requestId: %s", requestId)
	return nil
}

func resourceSslCertificateDelete(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("resource.cdnetworks_ssl_certificate.delete")
	var diags diag.Diagnostics
	var certificateId, _ = strconv.ParseInt(data.Id(), 10, 64)
	var response *certicate.DeleteCertificateForTerraformResponse
	var requestId string
	var err error
	err = resource.RetryContext(context, time.Duration(2)*time.Minute, func() *resource.RetryError {
		requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseSslCertificateClient().DeleteCertificate(certificateId)
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
	log.Printf("resource.cdnetworks_ssl_certificate.delete success")
	log.Printf("requestId: %s", requestId)
	return nil
}
