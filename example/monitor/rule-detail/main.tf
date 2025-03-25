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

data "cdnw_monitor_realtime_rules_detail" "test-rule" {
  rule_name = "test_rule_name"
}

output "show-test-rule" {
  value = data.cdnw_monitor_realtime_rules_detail.test-rule
}