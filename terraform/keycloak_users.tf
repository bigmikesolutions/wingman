resource "keycloak_user" "admin" {
  count = var.keycloak_admin_enabled ? 1 : 0

  username = var.keycloak_admin_user
  first_name = "Wingman"
  last_name = "Bot"
  email = "admin@wingman.com"

  enabled = true

  realm_id = keycloak_realm.wingman.id

  attributes = {
    foo = "bar"
    multivalue = "value1##value2"
  }

  initial_password {
    value     = var.keycloak_admin_password
    temporary = var.keycloak_admin_password_temporary
  }
}

resource "keycloak_user_groups" "admin" {
  count = var.keycloak_admin_enabled ? 1 : 0

  realm_id = keycloak_realm.wingman.id
  user_id  = keycloak_user.admin[0].id
  group_ids = [
    keycloak_group.admin.id
  ]
}