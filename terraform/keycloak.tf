resource "keycloak_realm" "wingman" {
  realm = var.keycloak_realm
}

resource "keycloak_openid_client" "wingman" {
  realm_id  = keycloak_realm.wingman.id

  client_id = var.wingman_client_id
  client_secret = var.wingman_client_secret

  name      = "Wingman"
  enabled   = true

  access_type = "PUBLIC"
  access_token_lifespan = var.wingman_access_token_lifespan

  service_accounts_enabled = false      # required for confidential access type
  standard_flow_enabled = true          # required for Authorization Code Flow
  direct_access_grants_enabled = true   # required for password grant
  implicit_flow_enabled = false

  valid_redirect_uris = [
    "http://traefik-auth:4181/oauth2/callback",
    "http://localhost:8088/oauth2/callback",
    "http://wingman/oauth2/callback",
  ]

  web_origins = [
    "*"
  ]

  login_theme = "keycloak"

  extra_config = {
    "key1" = "value1"
    "key2" = "value2"
  }
}

resource "keycloak_role" "wingman_api_read" {
  realm_id  = keycloak_realm.wingman.id
  name      = "api_read"
  client_id = keycloak_openid_client.wingman.id
}

resource "keycloak_role" "wingman_api_write" {
  realm_id  = keycloak_realm.wingman.id
  name      = "write_read"
  client_id = keycloak_openid_client.wingman.id
}

resource "keycloak_openid_client_scope" "wingman_scope" {
  realm_id               = keycloak_realm.wingman.id
  name                   = "wingman_scope"
  description            = "When requested, this scope will map a user's group memberships to a claim"
  include_in_token_scope = true
  gui_order              = 1
}

resource "keycloak_openid_client_optional_scopes" "wingman_optional_scopes" {
  realm_id  = keycloak_realm.wingman.id
  client_id = keycloak_openid_client.wingman.id

  optional_scopes = [
    "address",
    "phone",
    "offline_access",
    "microprofile-jwt",
    keycloak_openid_client_scope.wingman_scope.name
  ]
}

resource "keycloak_group" "admin" {
  realm_id = keycloak_realm.wingman.id
  name     = "Admins"
}

resource "keycloak_group_roles" "admin_roles" {
  realm_id = keycloak_realm.wingman.id
  group_id = keycloak_group.admin.id
  role_ids = [
    keycloak_role.wingman_api_read.id,
    keycloak_role.wingman_api_write.id,
  ]
}

