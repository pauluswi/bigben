package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pauluswi/bigben/exception"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLDatabase(configuration Config) *sql.DB {
	ctx, cancel := NewMySQLContext()
	defer cancel()

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true",
			configuration.Get("MYSQL_USER"),
			configuration.Get("MYSQL_PASSWORD"),
			configuration.Get("MYSQL_HOST"),
			configuration.Get("MYSQL_DB_NAME"),
		),
	)

	if err != nil {
		exception.PanicIfNeeded(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func NewMySQLContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
