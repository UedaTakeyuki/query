package query

import (
	"testing"
)

var q Query

func Test_01(t *testing.T) {
	q.SetTableName("tests").Select([]string{})
	if qs := q.GetQuery(); qs != "" {
		t.Errorf("query: %s\n", qs)
	}
}
