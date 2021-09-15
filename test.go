package query

import (
	"testing"

	"github.com/UedaTakeyuki/query"
)

var q query.Query

func Test_01(t *testing.T) {
	q.SetTableName("tests").Select([]string{})
	if qs := q.GetGetQuery; qs != "" {
		t.Error("query: %s\n", qs)
	}
}
