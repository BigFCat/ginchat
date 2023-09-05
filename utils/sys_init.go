package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var opens = map[string]func(string) gorm.Dialector{
	"mysql": mysql.Open,
	// "postgres":  postgres.Open,
	// "sqlite3":   sqlite.Open,
	// "sqlserver": sqlserver.Open,
}

func InitConfig() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("viper read config: ", viper.Get("settings"))
}

// 创建数据库连接池
func InitMySQL() *gorm.DB {
	newLoger:= logger.New(
		log.New(os.Stdout, "\r\n",log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel: logger.Info,
			Colorful: true,
		},
	)


	driverName := viper.GetString("settings.database.driver")
	host := viper.GetString("settings.database.host")
	port := viper.GetString("settings.database.port")
	dbname := viper.GetString("settings.database.dbname")
	username := viper.GetString("settings.database.username")
	password := viper.GetString("settings.database.password")
	charset := viper.GetString("settings.database.charset")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbname,
		charset,
	)

	db, err := gorm.Open(
		opens[driverName](args), 
		&gorm.Config{Logger:newLoger},
	)
	if err != nil {
		fmt.Println("fail err mysql", err.Error())
	}
	DB = db
	return db

}
