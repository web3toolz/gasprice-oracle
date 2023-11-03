package utils

import (
	"math/big"
	"sort"
)

func SortBigIntSlice(slice []*big.Int) {
	sort.Slice(slice, func(i, j int) bool { return slice[i].Cmp(slice[j]) == 1 })
}

func SortFloat64Slice(slice []float64) {
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
}

func SortInt64Slice(slice []int64) {
	sort.Slice(slice, func(i, j int) bool { return slice[i] > slice[j] })
}
