# Getting started 

## Setup environment variables 

`touch thenullchannel-dev/backend/api/.env`


### Populate the file with the following data
``` txt
DB_URI="host=postgres user=postgres password=pgadmin962 dbname=ideas port=5432 sslmode=disable"
ADMIN_USER=<inserts admin username here>
ADMIN_PASS=<inserts admin password here>
```

### Start the project
` docker-compose up -d ` 

the API will be accesible on http://localhost:8080/ideas and the frontend on http://localhost:3000


### Add ideas to the API
``` shell
curl -u admin_user:admin_pass -X POST -H "Content-Type: application/json" \
    -d '{"Description":"Using traefik as an ingress controller"}' \
    localhost:8080/create | jq
```

Response 
``` json 
{
    "id": 1,
    "description": "Using traefik as an ingress controller",
    "votes": 0
}
```


### Voting for an idea

``` bash 
curl -X POST -d '{id: 1,description: "Using traefik as an ingress controller",votes: 0}' localhost:8080/vote | jq 
```
Response 

HTTP 200 Ok.

Verify vote 

` curl localhost:8080/ideas | jq `

Resopnse

``` json 
{
    "id": 1,
    "description": "Using traefik as an ingress controller",
    "votes": 1
}
```


### Deleting Ideas

``` shell
curl -u admin_user:admin_pass -X DELETE -H "Content-Type: application/json" \
    -d '{"Description":"Using traefik as an ingress controller"}' \
    localhost:8080/ideas 
```
Response 

HTTP 200 Ok.










