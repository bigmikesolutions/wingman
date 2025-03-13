#!/bin/bash

curl -v http://localhost:8084/probes/health \
  -L \
  -H "Host: wingman" \
  -H 'X-Forwarded-Host:traefik-auth:4181' \
  -H "Authorization: Bearer ${TOKEN}"