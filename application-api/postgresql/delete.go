package postgresql

import "context"

func (p *PostgreSQL) Delete(deleteQuery string) error {
	rows, err := p.connPool.Query(context.Background(), deleteQuery)
	if nil != err {
		return err
	}

	defer rows.Close()

	return nil
}
