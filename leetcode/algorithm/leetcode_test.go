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
func Test12IntToRoman(t *testing.T) {
	params := []int{1994, 58, 9}
	params1 := []string{}
	for i := range params {
		r := intToRoman(params[i])
		params1 = append(params1, r)
		fmt.Println(r)
	}

	for i := range params1 {
		fmt.Println(romanToInt(params1[i]))
	}
}
func Test15ThreeSum(t *testing.T) {
	//nums := []int{0, 14, -7, 2, 7, 11, -9, 11, -12, 6, -10, -8, 9, -3, 7, -6, 3, 4, 14, -10, -8, 5, -1, 6, 12, 9, 12, -11, -15, -5, 5, 0, -6, 13, 9, 9, 10, 7, 5, 13, -2, 11, -10, -15, -15, 4, -14, -4, -15, 7, -7, -15, -3, 8, -2, -13, -6, -5, -9, -14, 5, 12, 1, 6, 2, -12, -8, -11, 10, 13, -13, -14, 1, 14, 8, 1, -4, 9, 4, -12, -6, 13, 10, 6, 6, -7, 8, 7, 3, 7, 8, -15, -4, -14, -1, 13, -11, 6, -12, -15, 4, 12, 8, -10, 4, 1, -1, 7, -13, -12, 10, -4, 8, 6, 10, -1, 6, -5, 13, -13, 9, 6, -13, -10}
	nums := []int{0, 0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}
func Test16ThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
func Test18FourSum(t *testing.T) {
	fourSumFes := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	for i := range fourSumFes {
		fmt.Println(fourSumFes[i])
	}
}

func Test22generateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis1(4))
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
func Test29Divide(t *testing.T) {
	fmt.Println(divide(20, 2))
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
func Test31nextPermutation(t *testing.T) {
	param := []int{1, 3, 5, 4, 2}
	nextPermutation(param)
	fmt.Println(param)
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
func Test33search(t *testing.T) {
	fmt.Println(search([]int{3, 1, 2}, 1))
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
		solveSudoku3(tests[i])
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
func Test38countAndSay(t *testing.T) {
	fmt.Println(countAndSay(6))
	//countAndSayOnce("111221")
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
func Test43multiply(t *testing.T) {
	//fmt.Println(multiply("498828660196","840477629533"))
	fmt.Println(multiply("12", "21"))
}
func Test44IsMatch(t *testing.T) {
	params := [][]string{
		//{"aa", "*"},
		//{"cb", "?a"},
		{"adceb", "ad*eb"},
		//{"aa", "a"},
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
func Test48rotate(t *testing.T) {
	param := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(param)
	for i := range param {
		fmt.Println(param[i])
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
func Test60GetPermutation(t *testing.T) {
	fmt.Println(getPermutation(4, 8))
}

func Test68FullJustify(t *testing.T) {
	fullJustify([]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		20)
}
func Test76minWindow(t *testing.T) {
	fmt.Println(minWindow("acbbaca", "aba"))
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
func Test97isInterleave(t *testing.T) {
	fmt.Println(isInterleave1("ab", "a", "abc"))
}
func Test99recoverTree(t *testing.T) {
	param := &TreeNode{
		Val: 17,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
	}
	recoverTree(param)
}
func Test114PreorderTraversal(t *testing.T) {
	res := preorderTraversal1(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}})
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
func Test127ladderLength(t *testing.T) {
	beginWord := "a"
	endWord := "c"
	wordList := []string{"a", "b", "c"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
func Test140WordBreak(t *testing.T) {
	//res := wordBreak1("pineapplepenapple", []string{"apple","pen","applepen","pine","pineapple"})
	res := wordBreak1("app", []string{"app"})
	for i := range res {
		fmt.Println(res[i])
	}
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
func Test147insertionSortList(t *testing.T) {
	readListNode(insertionSortList(generateListNode([]int{-1, 5, 3, 4, 0})))
	//readListNode(generateListNode([]int{4,2,1,3}))
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
func Test406reconstructQueue(t *testing.T) {
	//reconstructQueue([][]int{{7,0}, {4,4}, {7,1}, {5,0}, {6,1}, {5,2}})
	reconstructQueue([][]int{{9, 0}, {7, 0}, {1, 9}, {3, 0}, {2, 7}, {5, 3}, {6, 0}, {3, 4}, {6, 2}, {5, 2}})
}
func Test452findMinArrowShots(t *testing.T) {
	fmt.Println(findMinArrowShots([][]int{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}))
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
func Test514findRotateSteps(t *testing.T) {
	fmt.Println(findRotateSteps1("cotmaijx", "oatjiioicitatajtijciocjcaaxaaatmctxamacaamjjx"))
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
func Test763PartitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
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
func Test1356sortByBits(t *testing.T) {
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}

func generateListNode(nums []int) *ListNode {
	root := &ListNode{nums[0], new(ListNode)}
	node := root.Next
	for i := 1; i < len(nums); i++ {
		node.Val = nums[i]
		if i == len(nums)-1 {
			node.Next = nil
		} else {
			node.Next = new(ListNode)
		}
		node = node.Next
	}
	return root
}

func readListNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, ",")
		node = node.Next
	}
}
