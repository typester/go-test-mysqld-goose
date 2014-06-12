package mysqltest

import (
	"fmt"
	"path/filepath"
	"testing"

	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
)

func TestNew(t *testing.T) {
	mysql, err := New(filepath.Join("..", "testdb"))
	if err != nil {
		t.Error(err)
	}
	defer mysql.Server.Stop()

	db, err := sql.Open("mymysql", mysql.Dns())
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	version := ""
	if err := db.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
		t.Error(err)
	}
	if version == "" {
		t.Errorf("version should be set")
	}

	tables := make([]string, 0)

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		t.Error(err)
	}

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			t.Error(err)
		}
		fmt.Println("table:", table)
		tables = append(tables, table)
	}

	if len(tables) != 2 {
		t.Errorf("unexpected tables: %d, expected 2", len(tables))
	}

}
