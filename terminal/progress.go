package main

import (
	"fmt"
	"strings"
	"time"
)

func renderbar(count, total int) {
	barwidth := 60
	done := int(float64(barwidth) * float64(count) / float64(total))

	fmt.Printf("Progress: \x1b[33m%3d%%\x1b[0m ", count*100/total)
	fmt.Printf("\x1b[%dm[%s%s]\x1b[0m", count,
		strings.Repeat("=", done),
		strings.Repeat("-", barwidth-done))
}

func main() {
	total := 50
	for i := 1; i <= total; i++ {
		//<ESC>表示ASCII“转义”字符，0x1B
		fmt.Print("\x1b7")   // 保存光标位置 保存光标和Attrs <ESC> 7
		fmt.Print("\x1b[2k") // 清空当前行的内容 擦除线<ESC> [2K
		renderbar(i, total)
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\x1b8") // 恢复光标位置 恢复光标和Attrs <ESC> 8
	}
	fmt.Print("\x1b[1B")
	time.Sleep(time.Second * 1)

	fmt.Print("Done!")
}