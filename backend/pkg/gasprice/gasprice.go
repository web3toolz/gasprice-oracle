package gasprice

import (
	"fmt"
	"gasprice-oracle/utils"
	"github.com/montanaflynn/stats"
)

type Distribution struct {
	P40 int64
	P60 int64
	P75 int64
	P95 int64
}

func findPercentile(data []int64, percentile float64) (int64, error) {
	length := len(data)
	if length == 0 {
		return 0, nil
	} else if length < 3 {
		return data[length-1], nil
	}
	var dataAsFloat64 []float64
	for _, v := range data {
		dataAsFloat64 = append(dataAsFloat64, float64(v))
	}
	result, err := stats.Float64Data(dataAsFloat64).Percentile(percentile)
	if err != nil {
		return 0, err
	}
	return int64(result), err
}

func DistributionFomSlice(data []int64) (*Distribution, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data slice")
	}

	utils.SortInt64Slice(data)

	p40, err := findPercentile(data, 40)
	if err != nil {
		return nil, err
	}
	p60, err := findPercentile(data, 60)
	if err != nil {
		return nil, err
	}
	p75, err := findPercentile(data, 75)
	if err != nil {
		return nil, err
	}
	p95, err := findPercentile(data, 95)
	if err != nil {
		return nil, err
	}
	return &Distribution{
		P40: p40,
		P60: p60,
		P75: p75,
		P95: p95,
	}, nil
}
