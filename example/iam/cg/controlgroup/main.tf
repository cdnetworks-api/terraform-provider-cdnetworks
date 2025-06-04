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

resource "cdnetworks_iam_controlgroup" "controlgroup" {
  controlgroup_name     = "tf-example"
  account_list = ["account1", "account2"]
  domain_list = ["1.com", "2.com"]
}

resource "cdnetworks_iam_controlgroup" "controlgroup1" {
  controlgroup_name = "nntest"
  domain_list = ["alextest3.cdnetworks.com"]

}