# Clean Architecture in Go

Project skeleton for Clean Architecture in Go

Inspired by [Go Clean Architecture](https://github.com/bxcodec/go-clean-arch) by [bxcodec](https://github.com/bxcodec)

## Usage

To run

```shell
go run main.go
```

To test

```
go test ./...
```

### ENV Variables

It is set using [Viper](https://github.com/spf13/viper), you can see the implementation on `infrastructure/configuration`

It will set config from `.env` if it's exist,
otherwise it will set from system's ENV variables

## TODO

- [x] env variables	
- [ ] response format
- [ ] graceful shutdown
- [ ] db timeouts


## LICENSE

MIT
