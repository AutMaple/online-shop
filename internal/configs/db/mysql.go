package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"online.shop.autmaple.com/internal/configs/log"
)

type mysqlOption struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

var mysqlOpts *mysqlOption
var mysqlDB *sql.DB

func init() {
	mysqlOpts = &mysqlOption{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "root",
		Database: "online_shop",
	}
	openMysql()
}
func mysqlConnLink() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqlOpts.User,
		mysqlOpts.Password,
		mysqlOpts.Host,
		mysqlOpts.Port,
		mysqlOpts.Database,
	)
}

func openMysql() {
	db, err := sql.Open("mysql", mysqlConnLink())
	if err != nil {
		log.Fatal(err, "Open mysql failed")
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err, "Ping mysql failed")
	}
	mysqlDB = db
	log.Info("Connect to mysql successful")
}

func GetMysqlDB() *sql.DB {
	return mysqlDB
}
