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