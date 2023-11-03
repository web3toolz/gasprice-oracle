package ethclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gasprice-oracle/internal/config"
	"gasprice-oracle/utils"
	"io"
	"net/http"
)

type EthClient struct {
	nodeUrl string
}

type TransactionRaw struct {
	Hash     string `json:"hash"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
}

type BlockRaw struct {
	Hash         string           `json:"hash"`
	Number       string           `json:"number"`
	Timestamp    string           `json:"timestamp"`
	Transactions []TransactionRaw `json:"transactions"`
}

func (b *BlockRaw) GetBlockNumberAsInt() int64 {
	number, success := utils.HexToDecimal(b.Number)

	if success == false {
		return 0
	}

	return number.Int64()
}

func (b *BlockRaw) GetTransactionsGasPrices() ([]int64, error) {
	result := make([]int64, 0)

	for _, tx := range b.Transactions {
		gp, success := utils.HexToDecimal(tx.GasPrice)
		if success == false {
			return nil, fmt.Errorf("failed to parse gas price")
		}
		result = append(result, gp.Int64())
	}

	return result, nil
}

type JsonRPCResponse struct {
	Block BlockRaw `json:"result"`
	Error string   `json:"error"`
}

func New(ctx context.Context, config config.Config) (*EthClient, error) {
	return &EthClient{nodeUrl: config.NodeUrl}, nil
}

func (e *EthClient) GetBlockData(ctx context.Context, blockTagOrNumber string) (*BlockRaw, error) {
	payload := []byte(fmt.Sprintf(`{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "eth_getBlockByNumber",
			"params": ["%s", true]
		}`, blockTagOrNumber))

	resp, err := http.Post(e.nodeUrl, "application/json", bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result JsonRPCResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.Error != "" || result.Block.Number == "" {
		return nil, fmt.Errorf("failed to get block data")
	}

	return &result.Block, nil
}

func (e *EthClient) GetPendingBlockData(ctx context.Context) (*BlockRaw, error) {
	block, err := e.GetBlockData(ctx, "pending")

	if err != nil {
		return nil, err
	}

	return block, nil
}

func (e *EthClient) GetLatestBlockData(ctx context.Context) (*BlockRaw, error) {
	block, err := e.GetBlockData(ctx, "latest")

	if err != nil {
		return nil, err
	}

	return block, nil
}
