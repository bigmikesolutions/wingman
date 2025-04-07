#!/bin/bash

curl -X POST "http://keycloak:8080/realms/wingman/protocol/openid-connect/token" \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "grant_type=urn:ietf:params:oauth:grant-type:device_code" \
     -d "client_id=wingman" \
     -d "device_code=${DEVICE_CODE}"
