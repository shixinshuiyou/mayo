package db

import (
	"fmt"

	"github.com/shixinshuiyou/mayo/tool/config"
	"github.com/shixinshuiyou/mayo/tool/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMap = map[string]*gorm.DB{}

type dbInfo struct {
	CardName string // 数据连接别名、用于在配置文件中区分
	UserName string
	PassWord string
	Protocol string
	Host     string
	Port     int
	DBName   string //数据库名称
	Charset  string
}

func getDBInfo(dbCard string) *dbInfo {
	return &dbInfo{
		CardName: dbCard,
		UserName: config.Conf.Get(dbCard, "username").String("root"),
		PassWord: config.Conf.Get(dbCard, "password").String("123456"),
		Protocol: config.Conf.Get(dbCard, "protocol").String("tcp"),
		Host:     config.Conf.Get(dbCard, "host").String("127.0.0.1"),
		Port:     config.Conf.Get(dbCard, "port").Int(6379),
		DBName:   config.Conf.Get(dbCard, "dbname").String("test"),
		Charset:  config.Conf.Get(dbCard, "charset").String("utf8"),
	}
}

func (db *dbInfo) dsn() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		db.UserName,
		db.PassWord,
		db.Protocol,
		db.Host,
		db.Port,
		db.DBName)
}

func (db *dbInfo) initDB() *gorm.DB {

	dbLink, err := gorm.Open(mysql.New(mysql.Config{
		DSN: db.dsn(),
	}), &gorm.Config{})
	if err != nil {
		log.Logger.Panicf("failed to connect database:dbCard(%s)", db.CardName)
	}
	dbMap[db.CardName] = dbLink
	return dbLink
}

func GetDB(dbCard string) *gorm.DB {
	if _, ok := dbMap[dbCard]; ok {
		return dbMap[dbCard]
	}
	return getDBInfo(dbCard).initDB()
}
