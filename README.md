# mini-lens-go

## protofile 生成
### gRPC用
```
protoc -I . \
  --go_out ./ \
  --go_opt paths=source_relative \
  --go-grpc_out ./ \
  --go-grpc_opt paths=source_relative \
  proto/health.proto
```

### gateway用
