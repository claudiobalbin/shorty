# shorty
Simple url shortener in Go

# build

```sh
make build
```

# run

```sh
./shorty
```

# local test

## Unit tests
```sh
go test ./tests/unit/
```

## Integration tests

Start local environment
```sh
make local-environment
```

```sh
go test ./tests/integration/
```

# todos

- ~~Change url hash engine to base62;~~
- ~~Add clean archtecture folder structure;~~
- ~~Add persistance (redis/couchbase/etc);~~
- ~~Add .env with env-sample;~~
- Add tests;
- Add docker for tests and publishing;
- Add badges to README header;
