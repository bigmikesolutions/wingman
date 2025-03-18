#!/bin/bash

curl -X POST "http://localhost:8080/realms/wingman/protocol/openid-connect/userinfo" \
  -v \
  -H "Authorization: Bearer ${TOKEN}"