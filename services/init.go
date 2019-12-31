package services

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mobi.4se.tech/config"
	"xorm.io/xorm"
)

var db *xorm.Engine = nil

func init() {
	//配置DB
	base, err := xorm.NewEngine("mysql", config.MYSQL)
	if err != nil {
		fmt.Println("NewEngine:", err)
	}

	err = base.Ping()
	if err != nil {
		fmt.Println("Ping:", err)
	}
	db = base
}