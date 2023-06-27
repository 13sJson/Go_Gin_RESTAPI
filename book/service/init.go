package service

import (
	"book/model"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// kome
var DbEngine *xorm.Engine

/*
	initはDB(MySQL)との接続をするためのメソッド
	順番として[起動->middleware.info()->service.init()]を起動し接続してると思う、、、
*/

func init() {
	driverName := "mysql"
	DsName := "dbusername:dbpassword$@tcp([127.0.0.1]:3306)/dbname?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
