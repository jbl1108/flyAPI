# Readme for flyAPI

## Generate proto buf

protoc --go_out=. --go_opt=paths=source_relative protobuf\keyvalue\v1\keyvaluestore.proto