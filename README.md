# hello-go-openapi-apifirst

## Build

```
./build.sh
```

## curl

```
# add
curl -H 'content-type:application/json' -X POST --data-binary '{"name":"john"}' localhost:8080/pets

# list
curl localhost:8080/pets

# query
curl localhost:8080/pets/1000

# delete
curl -X DELETE localhost:8080/pets/1000
```



