package query

import (
	"fmt"
)

func (query_ptr *Query) Select(columns []interface{}) *Query {
	// confirm tableName is not ""
	if query_ptr.tableName == "" {
		query_ptr.err_str += "Need to set tableName before setting verb; "
		return query_ptr
	}

	// set verb as "Select"
	query_ptr.verb = Select

	// make body
	verb := "SELECT"
	params := ""
	for _, column := range columns {
		if column != "" {
			if params == "" {
				params += ToLiteralValue(column)
			} else {
				params += fmt.Sprintf(` ,%s`, ToLiteralValue(column))
			}
		}
	}
	if params == "" {
		params = "*"
	}
	query_ptr.body = fmt.Sprintf(`%s %s FROM %s`, verb, params, query_ptr.tableName)
	return query_ptr
}

func (query_ptr *Query) QueryStringSelect() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}
	query += ";"
	return
}
