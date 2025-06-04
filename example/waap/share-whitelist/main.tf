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

resource "cdnetworks_waap_share_whitelist" "demo" {
  rule_name            = "tf_test_update"
  relation_domain_list = ["waap.demo.com"]
  description          = "terraform test update"

  conditions {
    path_conditions {
      match_type = "NOT_EQUAL"
      paths      = ["/p1", "/p2"]
    }
    uri_conditions {
      match_type = "NOT_EQUAL"
      uri        = ["/uri1", "/uri2"]
    }
    ua_conditions {
      match_type = "NOT_EQUAL"
      ua         = ["ua1", "ua2"]
    }
    referer_conditions {
      match_type = "NOT_EQUAL"
      referer    = ["re1", "re2"]
    }
    header_conditions {
      match_type = "NOT_EQUAL"
      key        = "h"
      value_list = ["h11", "h21"]
    }
  }
}

data "cdnetworks_waap_share_whitelists" "demo" {
  rule_name = cdnetworks_waap_share_whitelist.demo.rule_name
}