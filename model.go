package gar

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dunstack/gar/scanner"
	"github.com/dunstack/gar/schema"
)

type BaseModel = schema.BaseModel

func Model[M any]() model[M] {
	var m M

	return model[M]{
		table: schema.TableOf(m),
	}
}

type model[M any] struct {
	table schema.Table
}

func (q model[M]) All(db *sql.DB, ctx context.Context) ([]M, error) {
	rows, err := db.QueryContext(ctx, q.query(), q.args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := scanner.Scan(rows)
	if err != nil {
		return nil, err
	}

	var results []M
	for _, r := range records {
		fmt.Println(r)
	}

	return results, nil
}

func (q model[M]) query() string {
	return fmt.Sprintf("SELECT * FROM %s", q.table.Name())
}

func (q model[M]) args() []any {
	return []any{}
}
