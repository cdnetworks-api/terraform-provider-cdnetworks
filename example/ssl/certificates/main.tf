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

data "cdnetworks_ssl_certificates" "myCertList" {
  name = "test20240625"
}

output "certList" {
  value = data.cdnetworks_ssl_certificates.myCertList
}