#!/bin/bash

curl -v http://localhost:8084/probes/health \
  -i -L \
  -H "Authorization: Bearer ${TOKEN}"