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

resource "cdnetworks_waap_bot_scene_whitelist" "demo" {
  domain      = "waap.example.com"
  name        = "test"
  description = "desc"
  conditions {
    match_name       = "IP_IPS"
    match_type       = "EQUAL"
    match_key        = ""
    match_value_list = ["1.1.1.1"]
  }
  conditions {
    match_name       = "PATH"
    match_type       = "EQUAL"
    match_key        = ""
    match_value_list = ["/path/test"]
  }
}

data "cdnetworks_waap_bot_scene_whitelists" "demo" {
  domain_list = [cdnetworks_waap_bot_scene_whitelist.demo.domain]
}

output "bot_scene_whitelist" {
  value = data.cdnetworks_waap_bot_scene_whitelists.demo
}