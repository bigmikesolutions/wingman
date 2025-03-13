#!/bin/bash

curl -X POST "http://localhost:8080/realms/wingman/protocol/openid-connect/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=password" \
     -d "client_id=wingman" \
     -d "client_secret=secret-123" \
     -d "username=admin" \
     -d "password=pass" \
     -d "scope=profile email"


http://keycloak:8080/realms/wingman/protocol/openid-connect/auth?client_id=wingman&redirect_uri=%3A%2F%2F%2F_oauth&response_type=code&scope=openid+profile+email&state=d6fb10b1f4563d563c9373ca5a2c6eea%3Aoidc%3A%3A%2F%2F%2Fverify