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
func Test6Convert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
func Test7myAtoi(t *testing.T) {
	fmt.Println(myAtoi("9223372036854775808"))
}
func Test7ReverseInt(t *testing.T) {
	fmt.Println(reverseInt(-1230))
}
func Test10IsMatch(t *testing.T) {
	tests := [][]string{
		{"", "*"},
		//{"aa", "a*"},
		//{"ab", ".*"},
		//{"aabc", "c*a*b*"},
		//{"mississippi", "mis*is*p*."},
		//{"aaa", "a*aa"},
		//{"mississippp", "mis*is*ip*pp"},
		//{"abcdfwaeqw", ".*"},
		//{"aaa", "a*aaaa"},
		//{"aba", "a.a"},
		//{"mississippi", "mis*is*ip*."},
	}
	for i := range tests {
		fmt.Println(IsMatch(tests[i][0], tests[i][1]))
		//fmt.Println(IsMatch1(tests[i][0], tests[i][1]))
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
func Test24SwapPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(swapPairs(head))
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
func Test39CombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
func Test40CombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
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
func Test44IsMatch(t *testing.T) {
	params := [][]string{
		{"aa", "*"},
		{"cb", "?a"},
		{"adceb", "*a*b"},
		{"aa", "a"},
	}
	for i := range params {
		fmt.Println(isMatch(params[i][0], params[i][1]))
	}
}
func Test45Jump(t *testing.T) {
	//nums := []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}
	//nums := []int{2, 3, 1, 1, 4}
	//nums := []int{2, 1, 1, 1}
	//nums := []int{5, 1, 1}
	nums := []int{8, 2, 4, 4, 4, 9, 5, 2, 5, 8, 8, 0, 8, 6, 9, 1, 1, 6, 3, 5, 1, 2, 6, 6, 0, 4, 8, 6, 0, 3, 2, 8, 7, 6, 5, 1, 7, 0, 3, 4, 8, 3, 5, 9, 0, 4, 0, 1, 0, 5, 9, 2, 0, 7, 0, 2, 1, 0, 8, 2, 5, 1, 2, 3, 9, 7, 4, 7, 0, 0, 1, 8, 5, 6, 7, 5, 1, 9, 9, 3, 5, 0, 7, 5}
	//nums := []int{9, 7, 9, 4, 8, 1, 6, 1, 5, 6, 2, 1, 7, 9, 0}

	fmt.Println(jump1(nums))
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
func Test52SolveNQueens2(t *testing.T) {
	res := totalNQueens(4)
	fmt.Println(res)
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
func Test77Combine(t *testing.T) {
	params := [][]int{
		{4, 2},
		{5, 5},
		{5, 3},
	}
	for _, param := range params {
		res := combine(param[0], param[1])
		for i := range res {
			fmt.Println(res[i])
		}
		fmt.Println("====")
	}

}
func Test94InorderTraversal(t *testing.T) {
	res := inorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}})
	fmt.Println(res)
}

func Test116Connect(t *testing.T) {
	node := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   4,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   5,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val:   6,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Right: &Node{
				Val:   7,
				Left:  nil,
				Right: nil,
				Next:  nil,
			},
			Next: nil,
		},
		Next: nil,
	}
	connect(node)
}
func Test143ReorderList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(head)
}
func Test216CombinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 9))
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

func Test347TopKFrequent(t *testing.T) {
	fmt.Println(topKFrequent1([]int{3, 0, 1, 0}, 1))
}
func Test377CombinationSum4(t *testing.T) {
	fmt.Println(combinationSum4_1([]int{1, 50}, 200))
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
func Test494FindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWays1([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
func Test530GetMinimumDifference(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(getMinimumDifference1(root))
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
func Test977SortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
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

func Test1002CommonChars(t *testing.T) {
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}
