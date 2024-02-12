package postgresql

import (
	"application-api/cfg"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func (p *PostgreSQL) Init() error {
	p.conf = Config{}

	err := cfg.Read(&p.conf, "./configs/postgresql.yml")
	if err != nil {
		return fmt.Errorf("PostgreSQL Config error : %s\n", err)
	}

	connectionDSN := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s pool_max_conns=%d statement_cache_capacity=%d",
		p.conf.User,
		p.conf.Password,
		p.conf.Host,
		p.conf.Port,
		p.conf.Database,
		p.conf.MaxConnections,
		p.conf.StatementCacheCapacity,
	)

	p.connPool, err = pgxpool.Connect(context.Background(), connectionDSN)

	return err
}
