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


docker run --rm \
-v ./migrations/sql:/flyway/sql \
-e FLYWAY_URL=jdbc:postgresql://localhost:5432/wingman \
-e FLYWAY_USER=admin \
-e FLYWAY_PASSWORD=pass \
flyway/flyway:11 migrate