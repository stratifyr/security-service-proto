# security-service-proto

```bash
    protoc \
      --proto_path=./proto \
      --go_out=./go/pb \
      --go_opt=paths=source_relative \
      --go-grpc_out=./go/pb \
      --go-grpc_opt=paths=source_relative \
      security-service.proto
```