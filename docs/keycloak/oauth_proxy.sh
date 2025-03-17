#!/bin/bash

curl -v http://localhost:8084/probes/health \
  -L \
  -H "Host: traefik:8080" \
  -H "Authorization: Bearer ${TOKEN}"