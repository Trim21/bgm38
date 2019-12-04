package model

import (
	"bgm38/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"time"
)

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s@(%s)/bgm38?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlAuth, config.MysqlHost))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(&Product{}, &Vote{}, &VoteOption{})
}
