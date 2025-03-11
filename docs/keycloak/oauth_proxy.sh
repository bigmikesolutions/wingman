#!/bin/bash

curl -v http://localhost:4180/oauth2/auth \
  -H "Authorization: Bearer ${TOKEN}"