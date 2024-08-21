package models

import (
	"os"
	"strconv"
	"time"

	"github.com/Cesta-UESC/cesta-backend/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	mysqlConnURL, err := utils.ConnectionURLBuilder("mysql")
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(mysqlConnURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	mysqldb, err := db.DB()
	if err != nil {
		return nil, err
	}

	mysqldb.SetMaxOpenConns(maxConn)
	mysqldb.SetMaxIdleConns(maxIdleConn)
	mysqldb.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	return db, nil
}
