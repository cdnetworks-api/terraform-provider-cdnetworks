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

resource "cdnetworks_waap_pre_deploy_waf" "demo" {
  domain        = "waap.example.com"
  config_switch = "ON"

  conf_basic {
    defend_mode      = "BLOCK"
    rule_update_mode = "AUTO"
  }

  rule_list {
    rule_id = 5002
    mode    = "BLOCK"
    exception_list {
      type         = "ip"
      match_type   = "EQUAL"
      content_list = ["192.168.1.1"]
    }
    exception_list {
      type         = "path"
      match_type   = "REGEX"
      content_list = ["/api/v1/.*"]
    }
  }

  rule_list {
    rule_id = 5003
    mode    = "LOG"
    exception_list {
      type         = "userAgent"
      match_type   = "CONTAIN"
      content_list = ["Mozilla"]
    }
  }
}

output "cdnetworks_waap_pre_deploy_result" {
  value = cdnetworks_waap_pre_deploy_waf.demo.host_list
}