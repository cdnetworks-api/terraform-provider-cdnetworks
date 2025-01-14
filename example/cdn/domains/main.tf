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

data "cdnetworks_cdn_domains" "myDomainList" {
  domain_names  = ["20240710001.conftest.com", "20240628003.conftest.com"]
  service_types = ["appa", "web"]
  page_size     = 1
  page_number   = 1
  status        = "enabled"
  start_time    = "2024-07-10T17:30:05+08:00"
  end_time      = "2024-07-10T18:31:05+08:00"
}

output "domain_list" {
  value = data.cdnetworks_cdn_domains.myDomainList
}