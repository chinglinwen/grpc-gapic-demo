# api demo

## run server

```
make build && ./demo 
```

## run grpc client

```
go run cmd/client/main.go 
2024/09/22 23:41:56 servers: [name:"s1" name:"s2"]
```

## run http client

```
$ curl  http://localhost:10000/v1/servers
{"servers":[{"name":"s1", "theme":""}, {"name":"s2", "theme":""}], "nextPageToken":""}
```