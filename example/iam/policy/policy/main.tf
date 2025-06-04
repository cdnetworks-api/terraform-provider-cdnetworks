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

resource "cdnetworks_iam_policy" "policy" {
  policy_name = "tf-example"
  description = "this is a policy test"
  policy_document = jsonencode([
    {
      "effect" : "allow",
      "action" : [
       "productCode:actionCode"
      ],
      "resource" : [
        "*"
      ]
    }
  ])
}

