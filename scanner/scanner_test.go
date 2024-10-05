package scanner_test

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/dunstack/gar/scanner"
	_ "github.com/mattn/go-sqlite3"
)

func TestScan(t *testing.T) {
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

	rows, err := db.Query("SELECT * FROM companies")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	got, err := scanner.Scan(rows)
	if err != nil {
		t.Fatal(err)
	}

	want := []scanner.H{
		{
			"id":      int64(1),
			"name":    "Paul",
			"age":     int64(32),
			"address": "California",
			"salary":  float64(20000),
		},
		{
			"id":      int64(2),
			"name":    "Allen",
			"age":     int64(25),
			"address": "Texas",
			"salary":  float64(15000),
		},
		{
			"id":      int64(3),
			"name":    "Teddy",
			"age":     int64(23),
			"address": "Norway",
			"salary":  float64(20000),
		},
		{
			"id":      int64(4),
			"name":    "Mark",
			"age":     int64(25),
			"address": "Rich-Mond",
			"salary":  float64(65000),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot : %+v\nwant: %+v", got, want)
	}
}
