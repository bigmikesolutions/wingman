resource "keycloak_group" "bms" {
  realm_id = keycloak_realm.wingman.id
  name     = "BigMikeSolutions"
}

resource "keycloak_role" "bms" {
  realm_id  = keycloak_realm.wingman.id
  name      = "bms"
  client_id = keycloak_openid_client.wingman.id
  description = "BMS organisation role"
}

resource "keycloak_group_roles" "bms_roles" {
  realm_id = keycloak_realm.wingman.id
  group_id = keycloak_group.bms.id
  role_ids = [
    keycloak_role.bms.id,
  ]
}