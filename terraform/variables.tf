variable "env" {
  description = "Environment (e.g., dev, staging, prod)"
  type        = string
}

variable "keycloak_realm" {
  description = "Keycloak realm"
  type        = string
  default     = "wingman"
}

variable "keycloak_endpoint" {
  description = "Keycloak server end-point"
  type        = string
}

variable "keycloak_cli_user" {
  description = "Keycloak CLI admin user name"
  type        = string
}

variable "keycloak_cli_password" {
  description = "Keycloak CLI admin password"
  type        = string
}

variable "keycloak_client_id" {
  description = "Keycloak client ID"
  type        = string
}

variable "keycloak_admin_enabled" {
  description = "Keycloak default global admin user creation enabled"
  type        = bool
  default     = false
}

variable "keycloak_admin_user" {
  description = "Keycloak default global admin user"
  type        = string
  default     = "admin"
}

variable "keycloak_admin_password" {
  description = "Keycloak default global admin user password"
  type        = string
}

variable "keycloak_admin_password_temporary" {
  description = "Make global admin password temporary"
  type        = bool
  default     = true
}

variable "wingman_client_id" {
  description = "Wingman client ID"
  type        = string
  default     = "wingman"
}

variable "wingman_client_secret" {
  description = "Wingman client secret"
  type        = string
}

variable "wingman_access_token_lifespan" {
  description = "Access token lifespan"
  type = number
  default = 3600
}