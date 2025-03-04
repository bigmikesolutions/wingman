terraform {
  required_providers {
    keycloak = {
      source = "mrparkers/keycloak"
      version = "4.4.0"
    }
  }
}

provider "keycloak" {
  url      = var.keycloak_endpoint
  username = var.keycloak_user
  password = var.keycloak_password
  client_id = var.keycloak_client_id
}