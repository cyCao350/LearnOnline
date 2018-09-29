package initiator

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var POSTGRES *gorm.DB

// postgres
//func init() {
//	//db init
//	connectString := config.GetPostGreConfig()
//	fmt.Println(connectString)
//	connect, err := gorm.Open(
//		"postgres",
//		connectString,
//	)
//	connect.LogMode(true)
//	if err != nil {
//		fmt.Println(err)
//		panic("connect postgres failed")
//	}
//	fmt.Println("Login postgres database success!")
//	POSTGRES = connect
//
//}


//mysql
func init() {
	//db init
	//connectString := config.GetPostGreConfig()
	//fmt.Println(connectString)
	connect, err := gorm.Open(
		"mysql",
		"root:1234567890@(127.0.0.1:3306)/ccydb?charset=utf8&parseTime=True&loc=Local",
	)
	connect.LogMode(true)
	if err != nil {
		fmt.Println(err)
		panic("connect postgres failed")
	}
	fmt.Println("Login postgres database success!")
	POSTGRES = connect

}

