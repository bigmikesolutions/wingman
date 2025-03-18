#!/bin/bash

curl -v http://localhost:8084/ \
  -i -L \
  -H "Authorization: Bearer ${TOKEN}"