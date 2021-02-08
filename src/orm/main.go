package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"orm/orm"
)

func main() {
	engine, _ := orm.NewEngine("mysql", "root:xuwenjin@tcp(127.0.0.1:3306)/data")
	defer engine.Close()
	se := engine.NewSession()
	_, _ = se.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = se.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := se.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
