package persistence

import (
	"fmt"
	"todo-sample/config"
	"todo-sample/infrastructure/persistence/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// データベース接続の処理を抽象化する
func NewDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.TodoModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
