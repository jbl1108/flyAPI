# Readme for flyAPI

## Generate proto buf

export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative protobuf\keyvalue\v1\keyvaluestore.proto

