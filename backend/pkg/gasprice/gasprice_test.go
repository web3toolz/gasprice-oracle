package gasprice

import (
	requirelib "github.com/stretchr/testify/require"
	"testing"
)

func TestFindPercentile(t *testing.T) {
	require := requirelib.New(t)

	normalSlice := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	shortSlice := []int64{0, 1}

	t.Run("should return 0 if no values", func(t *testing.T) {
		result, err := findPercentile([]int64{}, 0.5)
		require.NoError(err)
		require.Equal(int64(0), result)
	})

	t.Run("should return last value if short slice", func(t *testing.T) {
		result, err := findPercentile(shortSlice, 0.5)
		require.NoError(err)
		require.Equal(int64(1), result)
	})

	t.Run("should return 2", func(t *testing.T) {
		result, err := findPercentile(normalSlice, 50)

		require.NoError(err)
		require.Equal(int64(4), result)
	})

}
