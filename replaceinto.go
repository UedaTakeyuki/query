package query

import (
	"fmt"
)

func (query_ptr *Query) ReplaceInto(params []Param) *Query {

	// set verb as "ReplaceInto"
	query_ptr.verb = ReplaceInto

	// make body
	verb := "REPLACE INTO"

	return query_ptr.InsertReplaceCore(verb, params)
}

func (query_ptr *Query) QueryStringReplaceInto() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	/*	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}*/
	query += ";"
	return
}
