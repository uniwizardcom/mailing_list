package postgresql

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host                   string `yaml:"host"`
	Port                   int    `yaml:"port"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	Database               string `yaml:"database"`
	MaxConnections         int    `yaml:"maxConnections"`
	StatementCacheCapacity int    `yaml:"statementCacheCapacity"`
}

type PostgreSQL struct {
	conf     Config
	connPool *pgxpool.Pool

	// HandleMessage func(body []byte) error
}
