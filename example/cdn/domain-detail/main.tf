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

data "cdnetworks_cdn_domain_detail" "test-domain" {
  domain_name = "20240712001.conftest.com"
}

output "show-test-domain" {
  value = data.cdnetworks_cdn_domain_detail.test-domain
}