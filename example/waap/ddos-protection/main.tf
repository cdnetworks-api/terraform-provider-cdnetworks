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

data "cdnetworks_waap_ddos_protection_configs" "demo" {
  domain_list = ["example.waap.com"]
}

output "ddos_protection_configs" {
  value = data.cdnetworks_waap_ddos_protection_configs.demo
}