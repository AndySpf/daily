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

func TestMinPath(t *testing.T) {
}

func TestMinimalSteps(t *testing.T) {
	fmt.Println(minimalSteps([]string{".MM..", "#..M.", ".#..#", "..O..", ".S.OM", ".#M#T", "###..", "....."}))
}
