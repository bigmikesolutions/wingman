#!/bin/bash

open "http://localhost:8080/realms/wingman/protocol/openid-connect/auth?client_id=wingman&response_type=code&scope=openid&redirect_uri=http://localhost:8081/oauth2/callback"

