---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cdnetworks_cdn_domains Data Source - cdnetworks"
subcategory: "CDN"
description: |-
    Use this data source to query detailed information of CDN domain names.
---

# cdnetworks_cdn_domains (Data Source)

Use this data source to query detailed information of CDN domain names.

## Example Usage

```hcl
data "cdnetworks_cdn_domains" "myDomainList" {
  domain_names  = ["20240710001.conftest.com", "20240628003.conftest.com"]
  service_types = ["appa", "web"]
  page_size     = 1
  page_number   = 1
  status        = "enabled"
  start_time    = "2024-07-10T17:30:05+08:00"
  end_time      = "2024-07-10T18:31:05+08:00"
}

output "domain_list" {
  value = data.cdnetworks_cdn_domains.myDomainList
}
```




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domain_names` (List of String) Specify the accelerated domain name for the query. Multiple domain names are allowed. If not specified, all domain names will be searched by default.
- `end_time` (String) RFC3339 formatted date indicating the ending date. Example: 2024-01-01T22:30:00+08:00
- `page_number` (Number) Page number must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.
- `page_size` (Number) Number of domain name data items for paging, must be a positive integer greater than 0.If not passed, then no paging. If it is passed, pageSize is required.
- `service_types` (List of String) Specify the service type to be queried. Multiple services are allowed. Data will be returned if any one service is satisfied. If not passed, all services will be checked by default. For example: [wsa,waf], returns all domains whose services include wsa or include waf.
- `start_time` (String) RFC3339 formatted date indicating the starting date. Example: 2024-01-01T22:30:00+08:00
- `status` (String) Status of the accelerated domain. Optional value: enabled, disabled, deploying, checking, disabling, deployFailed, disableFailed.

### Read-Only

- `code` (String) Response code, 0 means successful.
- `data` (List of Object) Response data. (see [below for nested schema](#nestedatt--data))
- `id` (String) The ID of this resource.
- `message` (String) Response error message if failed.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `page_number` (Number)
- `page_size` (Number)
- `result_list` (List of Object) (see [below for nested schema](#nestedobjatt--data--result_list))
- `total_count` (Number)
- `total_page_number` (Number)

<a id="nestedobjatt--data--result_list"></a>
### Nested Schema for `data.result_list`

Read-Only:

- `cname` (String)
- `create_time` (String)
- `domain_id` (String)
- `domain_name` (String)
- `enabled` (String)
- `service_types` (List of String)
- `status` (String)
