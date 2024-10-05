package gar_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dunstack/gar"

	_ "github.com/mattn/go-sqlite3"
)

type Company struct {
	gar.BaseModel `gar:"table:companies"`

	ID int
}

func TestModel(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`
		CREATE TABLE companies (
			id				INT 		PRIMARY KEY     NOT NULL,
			name			TEXT    NOT NULL,
			age				INT     NOT NULL,
			address		CHAR(50),
			salary		REAL
		);

		INSERT INTO companies (id, name, age, address, salary) 
		VALUES	(1, 'Paul', 32, 'California', 20000.00),
						(2, 'Allen', 25, 'Texas', 15000.00),
						(3, 'Teddy', 23, 'Norway', 20000.00),
						(4, 'Mark', 25, 'Rich-Mond', 65000.00);
		;
	`); err != nil {
		t.Fatal(err)
	}

	t.Run("All", func(t *testing.T) {
		companies, err := gar.Model[Company]().All(db, context.Background())
		if err != nil {
			t.Error(err)
		}
		t.Errorf("%+v", companies)
	})
}
