/**
 * Created by BeastSanchez on 10/22/2017
 */

package shared

import "fmt"
import "github.com/jinzhu/gorm"

var db *gorm.DB
var err error

/*
dbType can be 'MySql', 'Postrges', ''
 */
func Init(dbType string) {
	db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/gingorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	//db.AutoMigrate(&models.Person{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}
