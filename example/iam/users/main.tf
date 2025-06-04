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

data "cdnetworks_iam_users" "user_list" {
  page_size     = 1
  page_number   = 10
}

output "user_list" {
  value = data.cdnetworks_iam_users.user_list
}