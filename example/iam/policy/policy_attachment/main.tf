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

resource "cdnetworks_iam_policy_attachment" "policyAttachment" {
  policy_name = ["policyName1","policyName2"]
  login_name = "subAccountLoginName"
}