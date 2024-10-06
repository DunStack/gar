package stmt

type BaseStmt struct {
	*BaseStmt

	Table string
	Where string
}

func (b BaseStmt) table() string {
	if b.Table != "" {
		return b.Table
	}
	return b.BaseStmt.table()
}

func (b BaseStmt) where() string {
	w []string
	w1 := b.BaseStmt.where()
	w2 := b.Where
	if 
}
