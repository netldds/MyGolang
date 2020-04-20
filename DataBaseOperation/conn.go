package DataBaseOperation

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func AA() {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"123",
		"localhost",
		3306,
		"taishan_dev2",
	)
	g, err := gorm.Open("mysql", dbUrl)
	g.DB().SetMaxOpenConns(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= 100000; i++ {
		_, err := g.DB().Conn(context.Background())
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v -> %v \n", i)
	}
	fmt.Println(os.Getwd())
	select {}
}
