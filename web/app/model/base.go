package model

import (
	"bgm38/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver

	"time"
)

type base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

//DB gorm database object
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s@(%s)/bgm38?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlAuth, config.MysqlHost))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Vote{}, &VoteOption{})
}
