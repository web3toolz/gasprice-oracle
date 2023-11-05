# Gasprice oracle

## Backend

### Install

```shell
cd backend
go install 
```

### Prepare configuration

Copy `config.example.yaml` to `config.yaml` to the same directory and edit it.

```shell
cd backend
cp config.example.yaml config.yaml
```

### Run server in development mode

```shell
cd backend
go run cmd/cli/main.go server
```

### Request data from server

```shell
curl -X GET http://localhost:8080/ | jq
```

### Build application

```shell
cd backend
go build -o gasprice-oracle cmd/cli/main.go
```