#!/bin/bash

curl -X POST "http://keycloak:8080/realms/wingman/protocol/openid-connect/auth/device" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "client_id=wingman"
