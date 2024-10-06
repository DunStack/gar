package stmt

import "fmt"

func Select(stmt BaseStmt) selectStmt {
	return selectStmt{
		BaseStmt: stmt,
	}
}

type selectStmt struct {
	BaseStmt
}

func (s selectStmt) Query() string {
	query := fmt.Sprintf("SELECT * FROM %s", s.table())
	if where := s.where(); where != "" {
		query = fmt.Sprintf("%s WHERE %s", query, where)
	}
	return query
}
