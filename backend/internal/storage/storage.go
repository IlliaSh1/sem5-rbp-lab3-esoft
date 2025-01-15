package storage

import (
	"fmt"

	"github.com/IlliaSh1/backend/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage(cfg *configs.DBConfig) (*Storage, error) {
	db, err := newDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Storage{
		DB: db,
	}, nil
}

func newDB(cfg *configs.DBConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var dsn string

	switch cfg.DBType {
	case "mysql":
		dsn = fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.AdminName,
			cfg.DBPassword,
			cfg.Host,
			cfg.Port,
			cfg.DBName,
		)
		fmt.Println(dsn)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("не удалось подключиться к базе данных, причина: %v", err)
		}
	default:
		return nil, fmt.Errorf("неизвестный тип базы данных")
	}

	return db, nil
}
