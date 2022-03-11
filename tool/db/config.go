package db

import (
	"fmt"

	"github.com/shixinshuiyou/mayo/tool/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultProto   = "tcp"
	DefaultCharset = "utf8"
)

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

func initDBCard(dbCard string) (db *dbInfo, err error) {
	err = config.Conf.Get(dbCard).Scan(db)
	if err != nil {
		return
	}

	db.CardName = dbCard
	if db.Protocol == "" {
		db.Protocol = DefaultProto
	}
	if db.Charset == "" {
		db.Charset = DefaultCharset
	}

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

func (db *dbInfo) initDB() (*gorm.DB, error) {

	// TODO  优化mysql日志输出
	return gorm.Open(mysql.New(mysql.Config{
		DSN: db.dsn(),
	}), &gorm.Config{})

}
