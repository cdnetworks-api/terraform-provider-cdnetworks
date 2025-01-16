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

resource "cdnetworks_waap_domain_copy" "demo" {
  source_domain  = "waap.demo.com"
  target_domains = ["waap.demo2.com", "waap.demo3.com"]
}