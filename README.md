# Gasprice oracle

## UI

### Install

```shell
cd u
yarn install
```

### Prepare configuration

Copy .env.template to .env.local and edit it.

```shell
cp .env.template .env.local
```

### Run server in development mode

Go to `https://localhost:3000` in your browser.

```shell
yarn dev
```

### Build static files

They will be placed in `out` directory.

```shell
yarn build
```

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