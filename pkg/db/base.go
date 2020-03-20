package db

import (
	"fmt"

	"bgm38/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

// type base struct {
//	ID        uint       `gorm:"primary_key" json:"id"`
//	CreatedAt time.Time  `json:"-"`
//	UpdatedAt time.Time  `json:"-"`
//	DeletedAt *time.Time `sql:"index" json:"-"`
// }

// Mysql gorm database object
var Mysql *gorm.DB
var MysqlX *sqlx.DB

func InitDB() {
	var err error
	var dsn = fmt.Sprintf("%s@(%s)/bgm_ip_viewer?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlAuth, config.MysqlHost)

	Mysql, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	Mysql.SingularTable(true)
	// Mysql.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Vote{}, &VoteOption{})

	MysqlX, err = sqlx.Connect("mysql", dsn)

}
