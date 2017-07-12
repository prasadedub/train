package main

import (
	"bytes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbConn gorm.DB

func init() {

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprint(user, ":", secret, "@tcp(", dbip, ":", dbport, ")/", dbschema))
	log.Debug("\n MySQL Database Connection String :", buffer.String())
	dbURL := buffer.String()
	dbConn, err = gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	dbConn.DB()
	dbConn.DB().Ping()
	dbConn.DB().SetMaxIdleConns(10)
	dbConn.DB().SetMaxOpenConns(20)
	dbConn.SingularTable(true)
	dbConn.LogMode(true)
	return

}

type Customer struct {
	ID     string `gorm:"column:ID"  json:"ID"`
	Name   string `gorm:"column:NAME"  json:"name"`
	Email  string `gorm:"column:EMAIL"  json:"email"`
	Status string `gorm:"column:STATUS"  json:"status"`
}

func main() {

	fmt.Println("Hello, playground")

	c := []Customer{}
	dbConn.Table("CUSTOMER").Where("STATUS = ?", "active").Scan(&c)

	fmt.Println("Active Customers list: ", c)
}
