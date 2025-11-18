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

data "cdnetworks_ssl_certificate_application_detail" "example" {
  order_id = "SO20251029194956687.199497441"
}

output "data" {
  value = data.cdnetworks_ssl_certificate_application_detail.example
}