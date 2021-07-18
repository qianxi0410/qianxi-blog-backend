package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	var dberr error
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/blog?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"),
		viper.GetString("mysql.host"), viper.GetInt("mysql.port"))

	db, dberr = gorm.Open(mysql.New(mysql.Config{
		DSN:                      DSN,
		DefaultStringSize:        1024,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,
	}), &gorm.Config{})

	if dberr != nil {
		log.Fatalf("connect to mysql failed ! %v", dberr)
	}
}
