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

data "cdnetworks_ssl_certificate_detail" "myCert" {
  certificate_id = "1464893"
}

output "data" {
  value = data.cdnetworks_ssl_certificate_detail.myCert
}