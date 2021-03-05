# protos

The protobufs for the cauldron project

```shell
# core
protoc \
    --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    core/data/v1alpha/*.proto \
    core/errors/v1alpha/*.proto \
    common/types/v1alpha/*.proto \
    cauldron/v1alpha/*.proto
```
