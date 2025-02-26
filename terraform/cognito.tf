resource "aws_cognito_user_pool" "test_pool" {
  name = var.cognito_user_pool_name
}

resource "aws_cognito_user_pool_client" "test_client" {
  name         =  var.cognito_app_client_name
  user_pool_id = aws_cognito_user_pool.test_pool.id
}

output "user_pool_id" {
  value = aws_cognito_user_pool.test_pool.id
}

output "app_client_id" {
  value = aws_cognito_user_pool_client.test_client.id
}