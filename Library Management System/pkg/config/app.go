package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Sreejith1962*@/simpleinterest?charset=utf8&parseTime=True&loc=Local")
	/*dsn := "root:Sreejith1962*@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	*/
	if err != nil {
		panic(err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}
