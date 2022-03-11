package db

import (
	"gorm.io/gorm"
)

var dbMap = map[string]*gorm.DB{}

func GetDB(dbCard string) (*gorm.DB, error) {
	if _, ok := dbMap[dbCard]; ok {
		return dbMap[dbCard], nil
	}

	dbInfo, err := initDBCard(dbCard)
	if err != nil {
		return nil, err
	}

	db, err := dbInfo.initDB()
	dbMap[dbCard] = db
	return db, err
}
