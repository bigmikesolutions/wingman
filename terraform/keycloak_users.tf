# resource "keycloak_user" "example_user" {
#   username = "example-user"
#   first_name = "Example"
#   last_name = "User"
#   email = "example@domain.com"
#   enabled = true
#   realm_id = keycloak_realm.wingman.id
#
#   attributes = {
#     foo = "bar"
#     multivalue = "value1##value2"
#   }
# }
#
# resource "keycloak_user_groups" "user_groups" {
#   realm_id = keycloak_realm.wingman.id
#   user_id = keycloak_user.example_user.id
#
#   group_ids  = [
#     keycloak_group.admins.id
#   ]
# }