package query

import (
	"testing"
)

var q Query

func Test_01(t *testing.T) {
	var qs string
	// SELECT *  FROM tests;
	q.SetTableName("tests").Select([]string{})
	if qs = q.QueryString(); qs != `SELECT * FROM tests;` {
		t.Errorf("query: %s\n", qs)
	}

	// SELECT * From tests WHERE ID = 1
	if qs = q.Where(Equal("ID", 1)).QueryString(); qs != `SELECT * FROM tests WHERE ID = 1;` {
		t.Errorf("query: %s\n", qs)
	}

	// ToLiteralValue(1)
	if literal := ToLiteralValue(1); literal != `1` {
		t.Errorf("literal: %s\n", literal)
	}

	// ToLiteralValue("name")
	if literal := ToLiteralValue("name"); literal != `"name"` {
		t.Errorf("literal: %s\n", literal)
	}

	// Drop Table tests;
	if qs = q.DropTable().QueryString(); qs != `DROP TABLE tests;` {
		t.Errorf("query: %s\n", qs)
	}

}
