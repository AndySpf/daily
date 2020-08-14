package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestChipInCoins(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 13
	fmt.Println(chipInCoins(coins, amount))
}
