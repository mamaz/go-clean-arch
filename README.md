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

## Packages

We're using 

- [chi](https://github.com/go-chi/chi) for API framework
- [sqlx](https://github.com/spf13/viper) for database
- [viper](https://github.com/spf13/viper) for env variables
- [testfify](https://github.com/stretchr/testify) for testing

### ENV

It is set using [Viper](https://github.com/spf13/viper), you can see the implementation on `infrastructure/configuration`

It will set config from `.env` if it's exist,
otherwise it will set from system's ENV

## TODO

- [x] env variables	
- [x] response format
- [ ] mocking
- [ ] graceful shutdown
- [ ] db timeouts


## LICENSE

MIT
