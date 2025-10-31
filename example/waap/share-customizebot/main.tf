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

resource "cdnetworks_waap_share_customizebot" "demo" {
  bot_name         = "test"
  bot_act          = "LOG"
  bot_description  = "desc"
  rela_domain_list = ["waap.zhangb68.com"]
  condition_list {
    condition_name       = "IP_IPS"
    condition_value_list = ["1.1.1.1"]
    condition_func       = "EQUAL"
    condition_key        = ""
  }
}


data "cdnetworks_waap_share_customizebots" "demo" {
  bot_name = cdnetworks_waap_share_customizebot.demo.bot_name
}

output "customizebot_list" {
  value = data.cdnetworks_waap_share_customizebots.demo
}