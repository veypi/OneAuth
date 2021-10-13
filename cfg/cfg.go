package cfg

import (
	"fmt"
	"github.com/veypi/utils/cmd"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Path = cmd.GetCfgPath("OneAuth", "oa")

var CFG = &struct {
	AdminUser      string
	Host           string
	LoggerPath     string
	LoggerLevel    string
	Key            string
	TimeFormat     string
	Debug          bool
	EXEDir         string
	EnableRegister bool
	DB             struct {
		Type string
		Addr string
		User string
		Pass string
		DB   string
	}
}{
	AdminUser:      "admin",
	Host:           "0.0.0.0:4001",
	LoggerPath:     "",
	LoggerLevel:    "debug",
	TimeFormat:     "2006/01/02 15:04:05",
	Debug:          true,
	EXEDir:         "./",
	EnableRegister: true,
	DB: struct {
		Type string
		Addr string
		User string
		Pass string
		DB   string
	}{
		//Type: "sqlite",
		Addr: "127.0.0.1:3306",
		//Addr: "oa.db",
		User: "root",
		Pass: "123456",
		DB:   "one_auth",
	},
}

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	if db == nil {
		ConnectDB()
	}
	return db
}
func ConnectDB() *gorm.DB {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", CFG.DB.User, CFG.DB.Pass, CFG.DB.Addr, CFG.DB.DB)
	if CFG.DB.Type == "sqlite" {
		conn = CFG.DB.Addr
		db, err = gorm.Open(sqlite.Open(conn), &gorm.Config{})
	} else {
		db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}
	return db
}
