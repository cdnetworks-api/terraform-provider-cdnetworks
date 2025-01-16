terraform {
  required_providers {
    cdnetworks = {
      source = "cdnetworks-api/cdnetworks"
    }
  }
}

provider "cdnetworks" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
}

resource "cdnetworks_ssl_certificate" "cert-example" {
  name = "cert-example-name"
  cert = var.cert2
  key  = var.key2
}