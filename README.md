# gbs-jwt

Development setup:

```sh
$ make tidy # module loading
$ make secrets && ./bin/gbs-jwt_secrets config/secrets/local.json
```

Regenerate protobuf:

```sh
$ make proto # ! REQUIRE protoc + gogoproto installed
```

Try in local version:
```
$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test0"}' -plaintext localhost:8082 grpc.Secrets/Create
$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test1"}' -plaintext localhost:8082 grpc.Secrets/Create
$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test2"}' -plaintext localhost:8082 grpc.Secrets/Create
$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test3"}' -plaintext localhost:8082 grpc.Secrets/Create

$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{}' -plaintext localhost:8082 grpc.Secrets/List
```


## questions
- Using 2 external (personal) libraries for go-grpc and go-log, ok or replace ?
