//
// db.go
// Copyright (C) 2024 veypi <i@veypi.com>
// 2024-09-20 16:10:16
// Distributed under terms of the MIT license.
//

package cfg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var cmdDB = CMD.SubCommand("db", "database operations")
var cmdMigrate = cmdDB.SubCommand("migrate", "migrate database")
var ObjList = make([]any, 0, 10)

func init() {
	cmdMigrate.Command = func() error {
		// create table without constraints
		DB().DisableForeignKeyConstraintWhenMigrating = true
		err := DB().AutoMigrate(ObjList...)
		if err != nil {
			return err
		}
		// create constraints
		DB().DisableForeignKeyConstraintWhenMigrating = false
		return DB().AutoMigrate(ObjList...)
	}
	cmdDB.SubCommand("drop", "drop database").Command = func() error {
		return DB().Migrator().DropTable(ObjList...)
	}
}

func DB() *gorm.DB {
	if db == nil {
		var err error
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: Config.DSN,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			panic(err)
		}
	}
	return db
}
