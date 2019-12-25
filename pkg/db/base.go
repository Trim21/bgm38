package db

import (
	"fmt"

	"bgm38/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver

)

// type base struct {
//	ID        uint       `gorm:"primary_key" json:"id"`
//	CreatedAt time.Time  `json:"-"`
//	UpdatedAt time.Time  `json:"-"`
//	DeletedAt *time.Time `sql:"index" json:"-"`
// }

// Mysql gorm database object
var Mysql *gorm.DB

func InitDB() {
	var err error
	Mysql, err = gorm.Open("mysql",
		fmt.Sprintf("%s@(%s)/bgm_ip_viewer?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlAuth, config.MysqlHost))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	Mysql.SingularTable(true)
	// Mysql.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Vote{}, &VoteOption{})
}
