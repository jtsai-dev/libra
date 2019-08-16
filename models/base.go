/*
 * @Descripttion:
 * @version:
 * @Author: yyeiei
 * @Date: 2019-08-12 14:56:31
 * @LastEditors: yyeiei
 * @LastEditTime: 2019-08-16 16:43:03
 */
package models

import (
	"fmt"

	"libra/pkg"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var X *xorm.Engine

// SyncDataBase sync the struct to db
func SyncDataBase() {
	X.Sync2(new(Account))
	X.Sync2(new(WxAccount))
	// X.Sync2(new(Directory))
	// X.Sync2(new(Option))
	X.Sync2(new(Login))
	X.Sync2(new(Adjudication))
	X.Sync2(new(Node))
}

func Setup() {
	var (
		err                                           error
		dbType, dbName, user, password, host, charset string
	)
	dbSetting := pkg.Configs.Database

	dbType = dbSetting.Type
	dbName = dbSetting.Name
	user = dbSetting.User
	password = dbSetting.Password
	host = dbSetting.Host
	charset = dbSetting.Charset

	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		user,
		password,
		host,
		dbName,
		charset,
	)

	X, err = xorm.NewEngine(dbType, connect)
	if err != nil {
		log.Error(err)
	}

	if pkg.Configs.Server.RunMode == "debug" {
		X.ShowSQL(true)
	}
}

func CloseDB() {
	defer X.Close()
}
