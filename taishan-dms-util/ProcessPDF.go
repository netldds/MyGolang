package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

func main() {
	flag.Set("logtostderr", "true")
	host := flag.String("h", "127.0.0.1", "-h 127.0.0.1")
	passwd := flag.String("p", "123", "-p 123")
	dbName := flag.String("db", "taishan_dev2", "-db taishan")
	//PrefixPath:=flag.String("prepath","")
	flag.Parse()
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		*passwd,
		*host,
		3306,
		*dbName)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		glog.Warning(err)
		return
	}
	Do(db)
}
