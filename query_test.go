package query

import (
	"testing"
)

var q Query

func Test_01(t *testing.T) {
	// SELECT *  FROM tests;
	q.SetTableName("tests").Select([]string{})
	if qs := q.GetQuery(); qs != "SELECT * FROM tests;" {
		t.Errorf("query: %s\n", qs)
	}
}
