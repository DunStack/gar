package stmt_test

import (
	"testing"

	"github.com/dunstack/gar/stmt"
)

func TestSelect(t *testing.T) {
	stmtUsers := stmt.BaseStmt{
		Table: "users",
	}
	stmtJohnUsers := stmt.BaseStmt{
		BaseStmt: &stmtUsers,
		Where:    "first_name = 'John'",
	}
	stmtJohnDoeUsers := stmt.BaseStmt{
		BaseStmt: &stmtJohnUsers,
		Where:    "last_name = 'Doe'",
	}
	stmtSelectUsers := stmt.Select(stmtUsers)
	stmtSelectJohnUsers := stmt.Select(stmtJohnUsers)
	stmtSelectJohnDoeUsers := stmt.Select(stmtJohnDoeUsers)

	if got, want := stmtSelectUsers.Query(), "SELECT * FROM users"; want != got {
		t.Errorf("\nwant: %s\ngot : %s", want, got)
	}

	if got, want := stmtSelectJohnUsers.Query(), "SELECT * FROM users WHERE first_name = 'John'"; want != got {
		t.Errorf("\nwant: %s\ngot : %s", want, got)
	}

	if got, want := stmtSelectJohnDoeUsers.Query(), "SELECT * FROM users WHERE first_name = 'John' AND last_name = 'Doe'"; want != got {
		t.Errorf("\nwant: %s\ngot : %s", want, got)
	}
}
