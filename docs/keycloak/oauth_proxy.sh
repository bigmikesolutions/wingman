#!/bin/bash

curl -X POST -v http://localhost:8084/graphql \
  -i -L \
  -H "Authorization: Bearer ${TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"test": "data"}'
