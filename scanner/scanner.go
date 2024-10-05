package scanner

import "database/sql"

type H map[string]any

func Scan(rows *sql.Rows) ([]H, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	records, size := []H{}, len(cols)

	for rows.Next() {
		values := make([]any, size)
		for i := range size {
			values[i] = &values[i]
		}

		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		record := make(H, size)
		for i, col := range cols {
			record[col] = values[i]
		}

		records = append(records, record)
	}

	return records, nil
}
