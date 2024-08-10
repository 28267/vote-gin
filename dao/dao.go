package dao

import (
	"log"
	"time"

	"vote-gin/config"
	"vote-gin/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	// defer Db.Close()
	//测试连接
	err = Db.DB().Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"mysql connect error": Db.Error})
	}

	log.Println("Successfully connected to the database!")
	Db.DB().SetConnMaxIdleTime(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
