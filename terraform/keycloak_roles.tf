resource "keycloak_role" "admin_read" {
  realm_id  = keycloak_realm.wingman.id
  name      = "admin-read"
  client_id = keycloak_openid_client.wingman.id
  description = "Admin write right within a single organisation"
}

resource "keycloak_role" "admin_write" {
  realm_id  = keycloak_realm.wingman.id
  name      = "admin-write"
  client_id = keycloak_openid_client.wingman.id
  description = "Admin write right within a single organisation"
}

resource "keycloak_role" "manager_read" {
  realm_id  = keycloak_realm.wingman.id
  name      = "manager-read"
  client_id = keycloak_openid_client.wingman.id
  description = "Manager write right within a single organisation"
}

resource "keycloak_role" "manager_write" {
  realm_id  = keycloak_realm.wingman.id
  name      = "manager-write"
  client_id = keycloak_openid_client.wingman.id
  description = "Manager write right within a single organisation"
}

resource "keycloak_role" "developer_read" {
  realm_id  = keycloak_realm.wingman.id
  name      = "dev-read"
  client_id = keycloak_openid_client.wingman.id
  description = "Developer write right within a single organisation"
}

resource "keycloak_role" "developer_write" {
  realm_id  = keycloak_realm.wingman.id
  name      = "dev-write"
  client_id = keycloak_openid_client.wingman.id
  description = "Developer write right within a single organisation"
}