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

  oauth2_device_authorization_grant_enabled = true


  valid_redirect_uris = [
    "http://oauth2-proxy:4180/oauth2/callback",
    "http://localhost:8084/oauth2/callback",
    "http://localhost:8088/oauth2/callback",
    "http://traefik:8084/oauth2/callback",
    "http://traefik:8080/oauth2/callback",
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

resource "keycloak_openid_audience_protocol_mapper" "audience_mapper" {
  realm_id    = keycloak_realm.wingman.id
  client_id   = keycloak_openid_client.wingman.id
  name        = "audience-mapper"
  included_custom_audience = "wingman"
  add_to_id_token = true
  add_to_access_token = true
}
