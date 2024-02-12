package postgresql

import (
	"context"
)

func (p *PostgreSQL) Create(insertQuery string, dest interface{}) error {
	rows, err := p.connPool.Query(context.Background(), insertQuery)
	if nil != err {
		return err
	}

	defer rows.Close()

	if false == rows.Next() {
		return rows.Err()
	}

	err = rows.Scan(dest)
	if err != nil {
		return err
	}

	return nil
}
