# Wait for some condition until running command

## Installation
```bash
go get -v github.com/Napas/wait-for
go install -v github.com/Napas/wait-for/cmd/wait-for
```

## Usage
### Wait for successful http request
```bash
wait-for http -url https://google.com echo "Will be output when http request was successful"
```

### Wait for successful grpc request
```bash
wait-for grpc -url service:8080 -service ServiceName echo "Will be output when grpc request was successful"
```

### More options
```bash
wait-for -h
```