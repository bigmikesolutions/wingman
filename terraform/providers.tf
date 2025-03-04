terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  required_providers {
    keycloak = {
      source = "mrparkers/keycloak"
      version = "4.4.0"
    }
  }
}

provider "aws" {
  region                      = var.aws_region
  access_key                  = var.aws_access_key
  secret_key                  = var.aws_secret_key
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
  endpoints {
    cognitoidp = var.cognito_idp_endpoint
  }
}

provider "keycloak" {
  url      = var.keycloak_endpoint
  username = var.keycloak_user
  password = var.keycloak_password
  client_id = var.keycloak_client_id
}