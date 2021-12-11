package cfg

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"github.com/veypi/utils/cmd"
	"github.com/veypi/utils/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Path = cmd.GetCfgPath("oa", "settings")
var Ctx, Cancel = context.WithCancel(context.Background())

var CFG = &struct {
	AdminUser     string
	Host          string
	LoggerPath    string
	LoggerLevel   string
	APPUUID       string
	APPKey        string
	TimeFormat    string
	TimeZone      string
	Debug         bool
	FileUrlPrefix string
	FireDir       string
	DB            struct {
		Type string
		Addr string
		User string
		Pass string
		DB   string
	}
	ES *config.Config
}{
	APPUUID:       "jU5Jo5hM",
	APPKey:        "cB43wF94MLTksyBK",
	AdminUser:     "admin",
	Host:          "0.0.0.0:4001",
	LoggerPath:    "",
	LoggerLevel:   "debug",
	TimeFormat:    "2006/01/02 15:04:05",
	TimeZone:      "Asia/Shanghai",
	Debug:         true,
	FileUrlPrefix: "/file",
	FireDir:       "/Users/light/test/media/",
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
	ES: &config.Config{
		URL:         "http://127.0.0.1:9200",
		Index:       "",
		Username:    "",
		Password:    "",
		Shards:      0,
		Replicas:    0,
		Sniff:       nil,
		Healthcheck: nil,
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

var gormCfg = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: false, // 使用单数表名，启用该选项后，`User` 表将是`user`
		NoLowerCase:   true,
	},
}

func ConnectDB() *gorm.DB {
	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", CFG.DB.User, CFG.DB.Pass, CFG.DB.Addr, CFG.DB.DB)
	if CFG.DB.Type == "sqlite" {
		conn = CFG.DB.Addr
		db, err = gorm.Open(sqlite.Open(conn), gormCfg)
	} else {
		db, err = gorm.Open(mysql.Open(conn), gormCfg)
	}

	if err != nil {
		panic(err)
	}
	return db
}

var es *elastic.Client

func ES() *elastic.Client {
	if es == nil {
		ConnectES()
	}
	return es
}

func ConnectES() *elastic.Client {
	var err error
	es, err = elastic.NewClientFromConfig(CFG.ES)
	if err != nil {
		log.Warn().Msgf("connect es failed: %s", err)
	}
	_, _, err = es.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		// Handle error
		log.Warn().Msgf("connect es failed: %s", err)
	}
	return es
}
