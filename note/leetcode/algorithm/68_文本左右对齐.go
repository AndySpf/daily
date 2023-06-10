package algorithm

import (
	"strings"
)

type line struct {
	startIndex int
	wordCount  int
	oriLength  int
}

func (l line) reset() line {
	l.startIndex = l.startIndex + l.wordCount
	l.wordCount = 0
	l.oriLength = 0
	return l
}

func fullJustify(words []string, maxWidth int) []string {
	// 先划分好每一行的数据
	lines := []line{}
	index := 0
	item := line{}
	for index < len(words) {
		if item.oriLength+len(words[index])+item.wordCount <= maxWidth {
			item.oriLength += len(words[index])
			item.wordCount++
			index++
		} else {
			lines = append(lines, item)
			item = item.reset()
		}

		if index == len(words) {
			lines = append(lines, item)
		}
	}

	//按照每一行数据的不同执行不同规则
	res := make([]string, len(lines))
	for i := range lines {
		if lines[i].wordCount == 1 { // 只有一个单词的左对齐，无论哪一行
			res[i] = repair(words[lines[i].startIndex], maxWidth)
			continue
		}

		if i == len(lines)-1 { // 最后一行，强制左对齐
			res[i] = repair(strings.Join(words[lines[i].startIndex:lines[i].startIndex+lines[i].wordCount], " "), maxWidth)
			break
		}

		// 通用规则
		defaultSpace := lines[i].wordCount - 1
		diff := maxWidth - defaultSpace - lines[i].oriLength
		switch {
		case diff == 0: // 不用补空格
			res[i] = strings.Join(words[lines[i].startIndex:lines[i].startIndex+lines[i].wordCount], " ")
		case diff%defaultSpace == 0: // 整数倍补空格
			count := diff / defaultSpace
			sep := " "
			for i := 0; i < count; i++ {
				sep += " "
			}
			res[i] = strings.Join(words[lines[i].startIndex:lines[i].startIndex+lines[i].wordCount], sep)
		default: // 非整数倍补充空格
			count := diff / defaultSpace
			sep := " "
			for i := 0; i < count; i++ {
				sep += " "
			}

			remainder := diff % defaultSpace
			for j := lines[i].startIndex; j < lines[i].startIndex+lines[i].wordCount; j++ {
				words[j] = words[j] + " "
				remainder--
				if remainder == 0 {
					break
				}
			}
			res[i] = strings.Join(words[lines[i].startIndex:lines[i].startIndex+lines[i].wordCount], sep)
		}
	}
	return res
}

func repair(s string, max int) string {
	count := max - len(s)
	res := ""
	for i := 0; i < count; i++ {
		res += " "
	}
	return s + res
}
