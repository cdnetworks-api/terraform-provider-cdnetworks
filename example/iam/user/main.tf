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

resource "cdnetworks_iam_user" "test_user" {
  login_name     = "tf-example"
  display_name   = "user_display_name"
  status         = "1"
  email          = "mail@example.com"
  console_enable = "1"
}