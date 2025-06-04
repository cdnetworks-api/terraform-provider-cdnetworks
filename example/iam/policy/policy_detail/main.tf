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

data "cdnetworks_iam_policy_detail" "test-policy" {
  policy_name = "policyName"
}

output "first_policy_name" {
  value = data.cdnetworks_iam_policy_detail.test-policy
}