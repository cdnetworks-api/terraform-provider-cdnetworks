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

data "cdnetworks_iam_user_detail" "user1" {
  login_name = "user1"
}

output "show-user1" {
  value = data.cdnetworks_iam_user_detail.user1
}