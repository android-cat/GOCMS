package database

import (
	"fmt"
	"go-cms/domain/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// dbコンテナが立ち上がるまで接続を行う
func RetryConnectDB(dialector gorm.Dialector, opt gorm.Option, count uint) error {
	var err error
	for count > 1 {
		if db, err = gorm.Open(dialector, opt); err != nil {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... coutn:%v\n", count)
			continue
		}
		break
	}
	return err
}

func NewDB() *gorm.DB {
	// MySQLに接続
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	if err := RetryConnectDB(mysql.Open(dsn), &gorm.Config{}, 100); err != nil {
		log.Fatalln(err)
	}

	// マイグレーションを実行
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
