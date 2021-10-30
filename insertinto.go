package query

import (
	"fmt"
)

func (query_ptr *Query) InsertInto(params []Param) *Query {

	// set verb as "InsertInto"
	query_ptr.verb = InsertInto

	// make body
	verb := "INSERT INTO"

	return query_ptr.InsertReplaceCore(verb, params)
}

func (query_ptr *Query) QueryStringInsertInto() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	/*	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}*/
	query += ";"
	return
}
