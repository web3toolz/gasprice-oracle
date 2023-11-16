# Gasprice oracle

<p align="left">
  <a href="https://github.com/web3toolz/gasprice-oracle/actions"><img alt="GitHub Workflow Status (with event)" src="https://img.shields.io/github/actions/workflow/status/web3toolz/gasprice-oracle/build_images_on_push.yaml"></a>
  &nbsp;
</p>

<hr/>
<h4><a target="_blank" href="https://gasprice.web3toolz.com/" rel="dofollow"><strong>Live Demo</strong></a></h4>
<hr/>

Tool is designed to empower Web3 developers by providing real-time, accurate, and relevant gas price data for the top EVM (Ethereum Virtual Machine) compatible blockchains. 
Our goal is to streamline your development process by offering easy access to essential network information through a free, public API.

### Features

* Top EVM Chains Support: Access gas price data from the most popular EVM-compatible blockchains.
* Real-Time Data: Stay up-to-date with the latest gas prices, ensuring efficient and cost-effective transactions. See the full list of supported chains below.
* Free Public API: Easy and open access for all developers. No API key required.

# Getting Started

## UI Component

**Install**

```shell
cd u
yarn install
```

**Prepare configuration**

Copy .env.template to .env.local and edit it.

```shell
cp .env.template .env.local
```

**Run server in development mode**

Go to `https://localhost:3000` in your browser.

```shell
yarn dev
```

**Build static files**

They will be placed in `out` directory.

```shell
yarn build
```

## Backend component

**Install**

```shell
cd backend
go install 
```

**Prepare configuration**

Copy `config.example.yaml` to `config.yaml` to the same directory and edit it.

```shell
cd backend
cp config.example.yaml config.yaml
```

**Run server in development mode**

```shell
cd backend
go run cmd/cli/main.go server
```

**Request data from server**

```shell
curl -X GET http://localhost:8080/ | jq
```

**Build application**

```shell
cd backend
go build -o gasprice-oracle cmd/cli/main.go
```


## API Documentation

### Endpoints

* `GET "/"` - get latest gas price data

### Response format

The API returns data in JSON format. Here is an example of a successful response:

```json
{
  "networkName": {
    "updatedAt": 1700149263,
    "slow": 100000000,
    "normal": 100000000,
    "fast": 100000000,
    "fastest": 100000000
  }
}
```

Slow, normal, fast, and fastest strategies for gas price are in wei.


## Supported Chains

* Ethereum
* Binance Smart Chain
* Polygon
* Avalanche
* Fantom
* Arbitrum One
* Base
* Fantom
* Optimism
* more to be added...

## License

Distributed under the MIT License. See LICENSE for more information.