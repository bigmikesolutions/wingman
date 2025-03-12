#!/bin/bash

curl -v http://localhost:4180/probes/health \
  -H "Authorization: Bearer ${TOKEN}"