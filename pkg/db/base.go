package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"bgm38/config"
)

var MysqlX *sqlx.DB

func InitDB() {
	var err error
	var dsn = fmt.Sprintf("%s@(%s)/bgm_ip_viewer?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlAuth, config.MysqlHost)
	MysqlX, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		logrus.Fatalln(err)
	}

}
