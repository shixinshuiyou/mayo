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

func getDBInfo(dbCard string) (db *dbInfo) {
	err := config.Conf.Get(dbCard).Scan(db)
	if err != nil {
		log.Logger.Panicf("init %s db conn error:%s", dbCard, err)
	}
	db.CardName = dbCard
	return
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
