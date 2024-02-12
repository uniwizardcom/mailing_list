package initializers

import (
	"application-api/postgresql"
	"os"
)

func PostgreSQL() *postgresql.PostgreSQL {
	pg := &postgresql.PostgreSQL{}

	err := pg.Init()
	if err != nil {
		os.Exit(1)
		return nil
	}

	return pg
}
