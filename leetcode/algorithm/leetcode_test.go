package algorithm

import (
	"fmt"
	"testing"
)

func Test5LongestPalindrome(t *testing.T) {
	tests := []string{
		"abba",
		"abad",
		"acbca",
		"daba",
		"ccc",
		"ababababa",
		"abccba",
		"abcba",
	}
	for i := range tests {
		fmt.Println(longestPalindrome1(tests[i]))
	}
}
func Test10IsMatch(t *testing.T) {
	tests := [][]string{
		{"aa", "a"},
		{"aa", "a*"},
		{"ab", ".*"},
		{"aabc", "c*a*b*"},
		{"mississippi", "mis*is*p*."},
		{"aaa", "a*aa"},
		{"mississippp", "mis*is*ip*pp"},
		{"abcdfwaeqw", ".*"},
		{"aaa", "a*aaaa"},
		{"aba", "a.a"},
		{"mississippi", "mis*is*ip*."},
	}
	for i := range tests {
		fmt.Print(IsMatch(tests[i][0], tests[i][1]))
		fmt.Println(IsMatch1(tests[i][0], tests[i][1]))
	}
}
func Test214ShortestPalindrome(t *testing.T) {
	tests := []string{
		"aacecaaa",
		"aaceecaaa",
		"",
		"a",
	}
	for i := range tests {
		fmt.Println(shortestPalindrome(tests[i]))
	}
}
func Test23MergeKLists(t *testing.T) {
	test := []*ListNode{
		{
			Val: -2,
			Next: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: -1,
					Next: &ListNode{
						Val:  -1,
						Next: nil,
					},
				},
			},
		},
		nil,
	}
	node := mergeKLists(test)
	for {
		fmt.Println(node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
}
func Test25ReverseKGroup(t *testing.T) {
	tests := []*ListNode{
		{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for i := range tests {
		//node := reverseKGroup(tests[i], 3)
		node := reverseKGroup1(tests[i], 5)
		for {
			fmt.Println(node)
			if node.Next == nil {
				break
			}
			node = node.Next
		}
	}

}
func Test30FindSubstring(t *testing.T) {
	tests := []map[string][]string{
		//{"barfoothefoobarman": []string{"foo", "bar"}},
		//{"wordgoodgoodgoodbest": []string{"word", "good", "best"}},
		//{"wordgoodgoodgoodbestword": []string{"word", "good", "best", "good"}},
		{"wordgoodgoodgoodbestword": []string{"word", "word", "best", "good"}},
	}
	for i := range tests {
		for p := range tests[i] {
			fmt.Println(findSubstring1(p, tests[i][p]))
			fmt.Println(findSubstring(p, tests[i][p]))
		}
	}
}
func Test32LongestValidParentheses(t *testing.T) {
	tests := []string{
		"(()())",
		"(())))()()())",
		"",
		"(",
		"()",
		"()(())",
	}
	for i := range tests {
		fmt.Println(longestValidParentheses(tests[i]))
	}
}
func Test37SolveSudoku(t *testing.T) {
	tests := [][][]byte{
		{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
	}
	for i := range tests {
		solveSudoku1(tests[i])
		for row := range tests[i] {
			for col := range tests[i][row] {
				fmt.Print(tests[i][row][col] - '0')
				fmt.Print("  ")
			}
			fmt.Printf("\n")
		}

		//fmt.Println(tests[i])
	}
}
func Test41FirstMissingPositive(t *testing.T) {
	tests := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{-1, 0, 3},
		{1, 2, 3},
	}
	for i := range tests {
		fmt.Println(firstMissingPositive(tests[i]))
		fmt.Println(firstMissingPositive1(tests[i]))
	}
}

func Test42TrappingRainWater(t *testing.T) {
	tests := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{0, 1, 0, 1, 1, 0, 1, 3, 2, 1, 2, 1},
	}
	for i := range tests {
		fmt.Println(trap(tests[i]))
		fmt.Println(trap1(tests[i]))
	}
}

func Test51SolveNQueens1(t *testing.T) {
	res := solveNQueens1(6)
	fmt.Println(len(res))
	//for i := range res {
	//	for j := range res[i] {
	//		fmt.Println(res[i][j])
	//	}
	//	fmt.Println("============")
	//}
	//[1 9 17 18 26 34]
	//[2 11 13 22 24 33]
	//[3 6 16 19 29 32]
	//[4 8 12 23 27 31]
}

func Test251BinaryTreePaths(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	res := binaryTreePaths1(root)
	for _, path := range res {
		fmt.Println(path)
	}
}

func Test494FindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays1([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}

func Test841CanVisitAllRooms(t *testing.T) {
	params := [][][]int{
		//{
		//	{1}, {2}, {3}, {},
		//},
		//{
		//	{1, 3}, {3, 0, 1}, {2}, {0},
		//},
		{
			{7, 16, 8, 16, 19, 8}, {10}, {9, 11}, {3, 14, 16, 19}, {8, 10, 19, 1, 7}, {13, 5, 10, 15, 4}, {6, 13}, {14, 14, 11, 12, 18}, {3}, {17, 9}, {1, 2, 6, 9, 6}, {12, 12, 2}, {4, 4}, {2, 13, 17}, {17}, {}, {11, 15}, {3, 5}, {15, 18, 5}, {7, 18, 1},
		},
	}
	for _, param := range params {
		fmt.Println(canVisitAllRooms(param))
	}
}

func Test486PredictTheWinner(t *testing.T) {
	params := [][]int{
		//{1, 5, 233, 7},
		//{1, 5, 2},
		//{1, 89, 3},
		{1, 3, 1},
		{1, 5, 233, 7},
		{1, 3, 3, 4, 10, 6},
	}
	for _, param := range params {
		fmt.Println(PredictTheWinner1(param))
	}
}

func TestOffer20IsNumber(t *testing.T) {
	params := []string{
		//"+100", "5e2", "-123", "3.1416", "-1E-16", "0123",
		//"12e", "1a3.14", "1.2.3", "+-5", "12e+5.4", "1", "1 ", " ", ".", "3.", "..",
		//"+.8", "+8.", "-.",
		"46.e3",
		".e3",
		"4e3.",
	}
	for i := range params {
		fmt.Println(isNumber(params[i]))
	}
}
