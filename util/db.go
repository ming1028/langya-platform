package util

import (
	"context"
	"database/sql"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetDB() *sql.DB {
	dbOnce.Do(func() {
		var err error
		db, err := sql.Open("mysql", viper.GetString("db.dsn"))
		if err != nil {
			panic(any("failed to connect mysql, error " + err.Error()))
		}
		db.SetMaxOpenConns(viper.GetInt("db.active"))
		db.SetMaxIdleConns(viper.GetInt("db.idle"))
		db.SetConnMaxLifetime(time.Minute)
	})
	return db
}

func Transaction(ctx context.Context, fc func(tx *sql.Tx) error) error {
	panicked := true
	tx, err := GetDB().Begin()
	if err != nil {
		return err
	}

	defer func() {
		if panicked || err != nil {
			_ = tx.Rollback()
		}
	}()

	err = fc(tx)
	if err == nil {
		err = tx.Commit()
	}

	panicked = false
	return err
}
