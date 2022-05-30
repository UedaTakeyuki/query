package query

import (
	"fmt"
	"log"
	"testing"

	"github.com/UedaTakeyuki/query"
)

var q query.Query

func Test_01(t *testing.T) {
	var qs, qs1, qs2 string

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

	json_func := query.JsonFunction{ // Json Functions is supported
		Body: fmt.Sprintf(
			`json_insert(attr, "$.%s", "%s", "$.%s", "%s")`, // In created SQL string, this json function string is treated as special
			"kero", // Fx: don't be quoted, even ordinaly string shoud be quoted.
			"kerokero",
			"keroyon",
			"bahahai",
		),
	}

	// name & value pair for update
	params := []query.Param{
		{Name: "name", Value: "frog"},
		{Name: "attr", Value: json_func},
	}

	// UPDATE without param
	if qs = q.Update(nil).Where(query.Equal("ID", 1)).QueryString(); qs != "UPDATE tests WHERE ID = 1;" {
		t.Errorf("query: %s\n", qs)
	}

	// UPDATE with param and SET after
	qs1 = q.Update(params).Where(query.Equal("ID", 1)).QueryString()
	qs2 = q.Update(nil).Set(params).Where(query.Equal("ID", 1)).QueryString()
	if qs1 != qs2 {
		t.Errorf("query1: %s\nquery2: %s\n", qs1, qs2)
	}
	log.Println(q.Update(params).Where(query.Equal("ID", 1)).QueryString())
	log.Println(q.Update(nil).Where(query.Equal("ID", 1)).QueryString())
	log.Println(q.Update(nil).Set(params).Where(query.Equal("ID", 1)).QueryString())

	// name & value pair for update
	params1 := []query.Param{
		{Name: "attr", Value: []map[string]interface{}{{"kerokero": 1}, {"kerokero": 2}}},
	}
	log.Println(q.Update(nil).Set(params1).Where(query.Equal("ID", 1)).QueryString())

	jsonParams := []query.JsonPathAndValue{
		{Path: "$.user", Value: map[string]interface{}{"name": "taro", "age": 10}},
		{Path: "$.point", Value: 11},
		{Path: "$.type", Value: "discount"},
		{Path: "$.expired", Value: true},
	}
	log.Println(q.Update(nil).SetJson_Set("attr", jsonParams).Where(query.Equal("ID", 1)).QueryString())
}
