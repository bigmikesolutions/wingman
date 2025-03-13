#!/bin/bash

curl -v http://localhost:8084/probes/health \
  -L \
  -H "Host: wingman" \
  -H "Authorization: Bearer ${TOKEN}"