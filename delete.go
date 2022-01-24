package query

import (
	"fmt"
)

func (query_ptr *Query) Delete() *Query {
	// confirm tableName is not ""
	if query_ptr.tableName == "" {
		query_ptr.err_str += "Need to set tableName before setting verb; "
		return query_ptr
	}

	// set verb as "Select"
	query_ptr.verb = Delete

	// make body
	verb := "DELETE"
	query_ptr.body = fmt.Sprintf(`%s FROM %s`, verb, query_ptr.tableName)
	return query_ptr
}

func (query_ptr *Query) QueryStringDelete() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}
	query += ";"
	return
}
