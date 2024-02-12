package postgresql

import (
	"context"
	"fmt"
)

func (p *PostgreSQL) Read(selectQuery string, dest interface{}) error {
	rows, err := p.connPool.Query(context.Background(), selectQuery)
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

func (p *PostgreSQL) ReadRecord(selectQuery string, dest *map[string]interface{}) error {
	rows, err := p.connPool.Query(context.Background(), selectQuery)
	if nil != err {
		return err
	}

	defer rows.Close()

	if false == rows.Next() {
		if rows.Err() == nil {
			return fmt.Errorf("Record not exist")
		}

		return rows.Err()
	}

	record, err := rows.Values()
	if nil != err {
		return err
	}

	fieldDesc := rows.FieldDescriptions()
	for i, c := range record {
		(*dest)[string(fieldDesc[i].Name)] = c
	}

	if nil != rows.Err() {
		return rows.Err()
	}

	return nil
}

func (p *PostgreSQL) ReadAll(selectQuery string, dest *[]map[string]interface{}) error {
	rows, err := p.connPool.Query(context.Background(), selectQuery)
	if nil != err {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		record, err := rows.Values()
		// err = rows.Scan(data...)
		if nil != err {
			return err
		}

		recordMap := make(map[string]interface{})
		fieldDesc := rows.FieldDescriptions()
		for i, c := range record {
			recordMap[string(fieldDesc[i].Name)] = c
			// fmt.Printf("Record[%d => %s] = [%s]\n", i, fieldDesc[i].Name, c)
		}

		*dest = append(*dest, recordMap)
	}

	if nil != rows.Err() {
		return rows.Err()
	}

	return nil
}
