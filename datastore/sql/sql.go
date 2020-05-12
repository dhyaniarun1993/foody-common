package sql

import (
	"database/sql"

	"github.com/opentracing/opentracing-go"
)

// Configuration provides configuration for SQL Driver
type Configuration struct {
	DSN                   string `required:"true"`
	MaxIdleConnections    int    `required:"true"`
	MaxOpenConnections    int    `required:"true"`
	ConnectionMaxLifetime string `required:"true"`
}

// CreatePool creates connection pool for SQL server
func CreatePool(configuration Configuration, driver string, tracer opentracing.Tracer) *DB {
	db, err := sql.Open(driver, configuration.DSN)
	if err != nil {
		panic(err)
	}

	return &DB{db, tracer}
}
