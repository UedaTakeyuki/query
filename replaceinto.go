package query

import (
	"fmt"
)

type Param struct {
	Name  string
	Value interface{}
}

func (query_ptr *Query) ReplaceInto(params []Param) *Query {

	// confirm tableName is not ""
	if query_ptr.tableName == "" {
		query_ptr.err_str += "Need to set tableName before setting verb; "
		return query_ptr
	}

	// confirm params
	if len(params) == 0 {
		query_ptr.err_str += "params of ReplaceInto should not blank array; "
		return query_ptr
	}

	// set verb as "ReplaceInto"
	query_ptr.verb = ReplaceInto

	// make body
	verb := "REPLACE INTO"

	columns := ""
	exprs := ""

	for _, param := range params {
		if columns == "" {
			columns += param.Name
			exprs += ToLiteralValue(param.Value)
		} else {
			columns += fmt.Sprintf(`,%s`, param.Name)
			exprs += fmt.Sprintf(`,%s`, ToLiteralValue(param.Value))
		}
	}

	query_ptr.body = fmt.Sprintf(`%s %s (%s) VALUES (%s)`, verb, query_ptr.tableName, columns, exprs)
	return query_ptr
}

func (query_ptr *Query) QueryStringReplaceInto() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	/*	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}*/
	query += ";"
	return
}
