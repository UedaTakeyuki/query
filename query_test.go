package query

import (
	"testing"

	"github.com/UedaTakeyuki/query"
)

var q query.Query

func Test_01(t *testing.T) {
	var qs string
	// SELECT *  FROM tests;
	s := make([]interface{}, 0)
	q.SetTableName("tests").Select(s)
	if qs = q.QueryString(); qs != `SELECT * FROM tests;` {
		t.Errorf("query: %s\n", qs)
	}

	// SELECT * From tests WHERE ID = 1
	if qs = q.Where(query.Equal("ID", 1)).QueryString(); qs != `SELECT * FROM tests WHERE ID = 1;` {
		t.Errorf("query: %s\n", qs)
	}

	// ToLiteralValue(1)
	if literal := query.ToLiteralValue(1); literal != `1` {
		t.Errorf("literal: %s\n", literal)
	}

	// ToLiteralValue('name')
	if literal := query.ToLiteralValue("name"); literal != `'name'` {
		t.Errorf("literal: %s\n", literal)
	}

	// Drop Table tests;
	if qs = q.DropTable().QueryString(); qs != `DROP TABLE tests;` {
		t.Errorf("query: %s\n", qs)
	}

	// DELETE FROM tests WHERE ID = 'kero';
	if qs = q.SetTableName("tests").Delete().Where(query.Equal("ID", "kero")).QueryString(); qs != `DELETE FROM tests WHERE ID = 'kero';` {
		t.Errorf("query: %s\n", qs)
	}

}
