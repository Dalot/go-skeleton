go-skeleton
-----------
A skeleton/boilerplate for a go web service

## Dependencies

```bash
make deps
```

## Run

### Local
```bash
go run main.go
```

#### Docker
```bash
docker run --rm -e ENV=dev -p8000:8000 go-skeleton
```

## Build

### Local
```bash
make build
```

### Docker

```bash
docker build -t go-skeleton .
```

## Test

```bash
make test
```
 
## Licence
See [LICENCE](LICENSE)
 
