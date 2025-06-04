terraform {
  required_providers {
    cdnetworks = {
      source = "registry.terraform.io/cdnetworks-api/cdnetworks"
    }
  }
}

provider "cdnetworks" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
}

data "cdnetworks_iam_controlgroup_detail" "controlgroup" {
  controlgroup_name = "controlgroup1"
}

output "show_controlgroup" {
  value = data.cdnetworks_iam_controlgroup_detail.controlgroup
}