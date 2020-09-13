package daos

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Estuary?charset=utf8&parseTime=True")
	fmt.Println("booking controller init")

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}
}
