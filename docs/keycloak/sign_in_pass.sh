#!/bin/bash

curl -X POST "http://localhost:8080/realms/wingman/protocol/openid-connect/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=password" \
     -d "client_id=wingman" \
     -d "client_secret=secret-123" \
     -d "username=admin" \
     -d "password=pass" \
     -d "scope=profile email"
