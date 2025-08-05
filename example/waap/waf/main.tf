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

data "cdnetworks_waap_waf_configs" "demo" {
  domain_list = ["waap.example.com"]
}

output "waf_configs" {
  value = data.cdnetworks_waap_waf_configs.demo
}