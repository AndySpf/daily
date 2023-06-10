package algorithm

import "strings"

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	direction := 1
	rowIndex := 0
	for i := range s {
		rows[rowIndex] += string(s[i])
		if rowIndex+direction == len(rows)-1 || rowIndex+direction == 0 {
			direction = ^direction + 1
		}
		rowIndex += direction
	}
	return strings.Join(rows, "")
}
