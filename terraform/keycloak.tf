resource "keycloak_realm" "wingman" {
  realm = "wingman"
}

resource "keycloak_role" "admin" {
  realm_id    = keycloak_realm.wingman.id
  name        = "Admin"
  description = "Wingman admin"
  attributes = {
    key = "value"
    multivalue = "value1##value2"
  }
}

resource "keycloak_group" "admins" {
  realm_id = keycloak_realm.wingman.id
  name     = "admins"
}

resource "keycloak_openid_client_scope" "openid_client_scope" {
  realm_id               = keycloak_realm.wingman.id
  name                   = "scope=1"
  description            = "When requested, this scope will map a user's group memberships to a claim"
  include_in_token_scope = true
  gui_order              = 1
}

resource "keycloak_openid_client" "wingman" {
  realm_id  = keycloak_realm.wingman.id
  client_id = "wingman"

  name      = "Wingman"
  enabled   = true

  access_type = "CONFIDENTIAL"
  valid_redirect_uris = [
    "http://localhost:8080/openid-callback"
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

resource "keycloak_openid_client_optional_scopes" "wingman_optional_scopes" {
  realm_id  = keycloak_realm.wingman.id
  client_id = keycloak_openid_client.wingman.id

  optional_scopes = [
    "address",
    "phone",
    "offline_access",
    "microprofile-jwt",
    keycloak_openid_client_scope.openid_client_scope.name
  ]
}
