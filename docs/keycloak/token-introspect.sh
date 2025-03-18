#!/bin/bash

curl -X POST "http://localhost:8080/realms/wingman/protocol/openid-connect/token/introspect" \
  -v \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "client_id=wingman" \
  -d "client_secret=secret-123" \
  -d "token=$TOKEN"