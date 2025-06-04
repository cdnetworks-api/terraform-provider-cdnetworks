terraform {
  required_providers {
    cdnetworks = {
      source = "cdnetworks-api/cdnetworks"
    }
  }
}

provider "cdnetworks" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
}

data "cdnetworks_iam_controlgroups" "controlgroupList" {
}

output "controlgroup_list" {
  value = data.cdnetworks_iam_controlgroups.controlgroupList
}