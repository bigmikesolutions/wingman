variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "keycloak_endpoint" {
  description = "Keycloak server end-point"
  type        = string
}

variable "keycloak_user" {
  description = "Keycloak admin user name"
  type        = string
}

variable "keycloak_password" {
  description = "Keycloak admin password"
  type        = string
}

variable "keycloak_client_id" {
  description = "Keycloak client ID"
  type        = string
}
