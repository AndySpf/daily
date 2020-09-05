package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Articles struct {
	Id    int64  `xorm:"pk autoincr"`
	Date  string `xorm:"date"`
	Title string `xorm:"title"`
	Href  string `xorm:"href"`
}

var DB *xorm.Engine

func Init() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/gocn_article?charset=utf8")
	if err != nil {
		panic("mysql初始化失败:" + err.Error())
	}
	err = engine.Sync2(new(Articles))
	if err != nil {
		panic(err.Error())
	}
	DB = engine
}
