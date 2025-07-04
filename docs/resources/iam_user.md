---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cdnetworks_iam_user Resource - cdnetworks"
subcategory: "IAM"
description: |-
  Use this resource to create and manage iam user.
---

# cdnetworks_iam_user (Resource)
Use this resource to create and manage iam user.

## Example Usage

```hcl
resource "cdnetworks_iam_user" "test_user" {
  login_name     = "tf-example"
  display_name   = "user_display_name"
  status   = "1"
  email    = "mail@example.com" 
  console_enable    = "1"      
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `login_name` (String) User login name

### Optional

- `console_enable` (Number) Whether console access is enabled. Options: 1-enabled, 0-disabled
- `display_name` (String) User display name
- `email` (String) User email address
- `status` (Number) User status. Options: 1-active, 0-inactive

### Read-Only

- `id` (String) The ID of this resource.
