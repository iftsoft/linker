# linker
Device linker for Self-Service Terminal project

Generate proto files
```shell
buf lint
buf generate
```

Build executables
```shell
go build -o manager-srv ./cmd/manager_server
go build -o callback-srv ./cmd/callback_server
go build -o manager-cli ./cmd/manager_client
go build -o callback-cli ./cmd/callback_client
```
