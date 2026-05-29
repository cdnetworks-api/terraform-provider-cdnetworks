# CLAUDE.md — terraform-provider-cdnetworks

## Overview

The official CDNetworks Terraform Provider, implemented in Go. It enables managing CDNetworks platform resources — CDN, SSL, WAAP, IAM, Monitor, and more — through Terraform IaC.

- **Go module**: `github.com/cdnetworks-api/terraform-provider-cdnetworks`
- **Go version**: 1.22
- **License**: Mozilla Public License 2.0

---

## Build & Run

```bash
# Build the provider
go build .

# Tidy dependencies
go mod tidy

# Enable local debug mode
export debuggable=true
export TF_LOG=DEBUG
```

To upgrade the SDK version: update the `github.com/cdnetworks-api/cdnetworks-sdk-go` version in `go.mod`, then run `go mod tidy`.

---

## Directory Structure

```
main.go                          # Provider entry point
cdnetworks/
  provider.go                    # Provider registration, schema definition, auth config
  connectivity/
    client.go                    # CdnetworksClient: lazy-initialized service SDK clients
  common/
    common.go                    # Shared utility functions
    helper.go                    # Hash utilities
    provider.go                  # ProviderMeta interface definition
    validators.go                # Schema validators
  services/                      # Resource implementations, organized by service
    cdn/domain/
    ssl/certificate/
    ssl/certificateapplication/
    waap/{whitelist,customizerule,ratelimit,domain,waf,bot,predeploy,share-*/}
    iam/{cgmanage,policy,user}/
    monitor/rule/
docs/                            # Terraform Registry docs (do not edit manually)
example/                         # Example Terraform configurations
```

Empty placeholder directories (not yet implemented): `cdnetworks/services/appa/`, `security/`, `vpc/`

---

## Architecture & Code Conventions

### Provider Authentication

```hcl
provider "cdnetworks" {
  secret_id  = "..."   # or env var CDNETWORKS_SECRET_ID
  secret_key = "..."   # or env var CDNETWORKS_SECRET_KEY
  # optional: protocol, domain, service_type, contract_id, item_id
}
```

### ProviderMeta Interface

All resources obtain the SDK client via the `meta` parameter:

```go
client := meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseXxxClient()
```

### Adding a New Service Client

1. Add a `xxxConn *xxx.Client` field to the `CdnetworksClient` struct in `cdnetworks/connectivity/client.go`
2. Add the corresponding `UseXxxClient()` lazy-init method (check for nil before initializing, following the existing pattern)

### File Naming Conventions

Each service directory typically contains:
- `resource_<service>_<name>.go` — CRUD resource implementation
- `datasource_<service>_<name>.go` — data source (list / detail)

### Standard CRUD Implementation Pattern

```go
func ResourceXxx() *schema.Resource {
    return &schema.Resource{
        CreateContext: resourceXxxCreate,
        ReadContext:   resourceXxxRead,
        UpdateContext: resourceXxxUpdate,
        DeleteContext: resourceXxxDelete,
        Schema: map[string]*schema.Schema{ ... },
    }
}

func resourceXxxCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
    log.Printf("resource.cdnetworks_xxx.create")
    var diags diag.Diagnostics
    request := &xxx.CreateXxxRequest{}
    // Populate request fields from d.Get() using pointers

    var response *xxx.CreateXxxResponse
    var requestId string
    var err error
    err = resource.RetryContext(ctx, 2*time.Minute, func() *resource.RetryError {
        requestId, response, err = meta.(cdnetworksCommon.ProviderMeta).GetAPIV3Conn().UseXxxClient().Method(request)
        if err != nil {
            return resource.NonRetryableError(err)
        }
        return nil
    })
    if err != nil {
        return append(diags, diag.FromErr(err)...)
    }
    if response == nil || response.Data == nil {
        d.SetId("")
        return nil
    }
    d.SetId(...)
    log.Printf("resource.cdnetworks_xxx.create success, requestId: %s", requestId)
    time.Sleep(2 * time.Second)  // Wait for API eventual consistency
    return resourceXxxRead(ctx, d, meta)
}
```

Key conventions:
- All SDK calls must be wrapped in `resource.RetryContext` with a 2-minute timeout
- Always use `resource.NonRetryableError` for errors (unless the API explicitly supports retries)
- After a successful Create/Update, call Read to refresh state; add `time.Sleep(2 * time.Second)` in between
- When response is nil, call `d.SetId("")` and return
- Field assignment: `_ = d.Set("field", value)` (ignore the error returned by Set)

### Registering Resources in provider.go

```go
// Register a resource in ResourcesMap
"cdnetworks_xxx_yyy": xxx.ResourceXxxYyy(),

// Register a data source in DataSourcesMap
"cdnetworks_xxx_yyy": xxx.DataSourceXxxYyy(),
```

---

## Resources & Data Sources Reference

| Service | Resources | Data Sources |
|---------|-----------|--------------|
| CDN | `cdnetworks_cdn_domain` | `cdnetworks_cdn_domains`, `..._detail` |
| SSL | `cdnetworks_ssl_certificate`, `..._application` | list + detail |
| WAAP | whitelist, customizerule, ratelimit, domain, domain_copy, waf_config, waf_rule_exception, bot_config, predeploy×5, share×3 | corresponding list |
| IAM | `cdnetworks_iam_controlgroup`, `..._policy`, `..._policy_attachment`, `..._user` | list + detail |
| Monitor | `cdnetworks_monitor_realtime_rule` | `..._detail` |

---

## Notes

- **No automated tests**: The project currently has no `*_test.go` files. Validate changes manually using the Terraform configurations in the `example/` directory.
- **docs/ directory**: Auto-generated by tooling — do not edit manually.
- **Proxy configuration**: In a corporate network environment, set the `http_proxy`/`https_proxy` environment variables accordingly.
