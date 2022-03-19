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
$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test3", "claims": {"claimtest": "value"}}' -plaintext localhost:8082 grpc.Secrets/Create

$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{}' -plaintext localhost:8082 grpc.Secrets/List

$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test3", "claims": {"claimtest": "valueother"}}' -plaintext localhost:8082 grpc.Secrets/Update

$ grpcurl -v -import-path ../../.. -proto cmd/secrets/grpc/secrets.proto -d '{"name": "test3"}' -plaintext localhost:8082 grpc.Secrets/Delete

```


## questions
- claims should be map[string]interface{} all along instead of map[string]string

