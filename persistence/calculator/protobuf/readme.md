# caution

## 1. protobuf

dont touch `*.pb.go` files.

this directory is generated by `protoc` command.

## 2. generate

```bash
$ protoc --go_out=. --go_opt=paths=source_relative *.proto
```