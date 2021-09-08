package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

var (
	DB *gorm.DB
)

func Mysql(hostname string, port int, username string, password string, dbname string) (*gorm.DB, error) {
	dbCoon := username + ":" + password + "@tcp(" + hostname + ":" + strconv.Itoa(port )+ ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbCoon)
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
