#!/bin/bash

curl -X POST "http://localhost:8082/realms/wingman/protocol/openid-connect/token" \
     -d "grant_type=client_credentials" \
     -d "client_id=wingman" \
     -d "client_secret=secret-123" \
     -d "scope=profile email"