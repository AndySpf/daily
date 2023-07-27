package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"path"
	"strconv"
	"strings"
)

var begin int

func main() {
	fmt.Println("请输入开始转换的户号：")
	fmt.Scanln(&begin)
	abs, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	basePath := abs + "/picture"
	fs, err := os.ReadDir(basePath)
	if err != nil {
		panic(err)
	}
	for _, dirs := range fs {
		if dirs.IsDir() {
			fmt.Println("一级", dirs.Name())
			number, err := strconv.Atoi(dirs.Name())
			if err != nil {
				fmt.Println("无效的目录名:" + dirs.Name())
			}
			if number >= begin {
				imgs, err := os.ReadDir(path.Join(basePath, dirs.Name()))
				if err != nil {
					panic(err)
				}
				for _, item := range imgs {
					fmt.Println("二级", item.Name())
					if !item.IsDir() && strings.HasSuffix(item.Name(), ".jpg") {
						src, err := imaging.Open(path.Join(basePath, dirs.Name(), item.Name()))
						if err != nil {
							panic(err)
						}
						dst := imaging.Rotate(src, 270, nil)
						err = imaging.Save(dst, path.Join(basePath, dirs.Name(), strings.Replace(item.Name(), ".jpg", ".png", -1)))
						if err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}

}
