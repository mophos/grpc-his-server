Running

```
USER=xxx PASS=xxxx go run server.go
```

Complied
```
protoc --go_out=. --go-grpc_out=. proto/*.proto
```