# wingman



#### CLI

```shell
wingman plugins list
wingman auth
wingman aws sqs {command} {params}
wingman aws sqs peek -name {queue_name}
wingman k8s pods get -n default 
```

#### Rest API

```shell
POST /k8s/auth/sign_in
GET /k8s/v1/namespaces/{namespace}/pods
GET /k8s/{namespace}/srv
GET /api/v1/pods?limit=500

-- select * from table_name 
GET /postgres/db/{db_name}/tables/{table_name}?id=5&limit=10
 
```

#### Keycloak

Login command for docker-compose:

```shell
curl -X POST "http://localhost:8081/realms/master/protocol/openid-connect/token" \
-H "Content-Type: application/x-www-form-urlencoded" \
-d "client_id=admin-cli" \
-d "username=admin" \
-d "password=pass" \
-d "grant_type=password"
```

```shell
docker exec -it wingman_server sh -c "wget -S -O - --header='X-Forwarded-Proto:http' --header='X-Forwarded-Host:traefik-auth:4181'  'http://traefik-auth:4181'"
```


#  traefik-auth:
#    image: thomseddon/traefik-forward-auth
#    container_name: wingman_traefik_auth
#    command:
#      - --log-level=debug
#      - --insecure-cookie
#      - --secret=secret-123
#      - --default-provider=oidc
#      - --auth-host="traefik-auth:4181"
#      - --url-path=/oauth2/callback
#      - --providers.oidc.issuer-url=http://keycloak:8080/realms/wingman
#      - --providers.oidc.client-id=wingman
#      - --providers.oidc.client-secret=secret-123
#      - --cookie-domain=wingman
##      - --default-action=allow
#    restart: unless-stopped
#    depends_on:
#      keycloak:
#        condition: service_healthy
#    ports:
#      - "4181:4181"
#    labels:
#      - "traefik.enable=true"
#      - "traefik.http.services.traefik-auth.loadbalancer.server.port=4181"
#      - "traefik.http.routers.traefik-auth.entrypoints=web"
#      - "traefik.http.routers.traefik-auth.rule=PathPrefix(`/oauth2`)"
#      - "traefik.http.middlewares.traefik-auth.forwardauth.address=http://traefik-auth:4181"
#      - "traefik.http.middlewares.traefik-auth.forwardauth.trustForwardHeader=true"
#      - "traefik.http.middlewares.traefik-auth.forwardauth.authResponseHeaders=X-Auth-User,X-Auth-Email,X-Forwarded-User,X-Forwarded-Proto,X-Forwarded-Host,Authorization"
