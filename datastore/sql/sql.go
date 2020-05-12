package sql

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/opentracing/opentracing-go"
)

// Configuration provides configuration for SQL Driver
type Configuration struct {
	DSN                   string `required:"true"`
	MaxIdleConnections    int    `required:"true" split_words:"true"`
	MaxOpenConnections    int    `required:"true" split_words:"true"`
	ConnectionMaxLifetime string `required:"true" split_words:"true"`
}

// CreatePool creates connection pool for SQL server
func CreatePool(configuration Configuration, driver string, tracer opentracing.Tracer) *DB {
	db, err := sql.Open(driver, configuration.DSN)
	if err != nil {
		panic(err)
	}

	return &DB{db, tracer}
}
