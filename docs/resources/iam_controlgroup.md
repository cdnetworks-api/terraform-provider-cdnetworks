---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cdnetworks_iam_controlgroup Resource - cdnetworks"
subcategory: "IAM"
description: |-
  Use this resource to create and manage control group.
---

# cdnetworks_iam_controlgroup (Resource)
Use this resource to create and manage control group.

## Example Usage

```hcl
resource "cdnetworks_iam_controlgroup" "controlgroup" {
  controlgroup_name     = "tf-example"
  account_list = ["account1", "account2"]
  domain_list = ["1.com", "2.com"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `controlgroup_name` (String) Control Group Name

### Optional

- `account_list` (List of String) Account object array, Used to specify accounts with permission.
- `domain_list` (List of String) Domain array, Used to specify the domain contained in the Control Group

### Read-Only

- `id` (String) The ID of this resource.
