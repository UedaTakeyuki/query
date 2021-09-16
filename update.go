package query

import (
	"fmt"
)

func (query_ptr *Query) Update(params []Param) *Query {

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
	query_ptr.verb = Update

	// make body
	verb := "UPDATE"

	set_body := ""

	for _, param := range params {
		if set_body == "" {
			set_body += fmt.Sprintf(`%s=%s`, param.Name, ToLiteralValue(param.Value))
		} else {
			set_body += fmt.Sprintf(`,%s=%s`, param.Name, ToLiteralValue(param.Value))
		}
	}

	query_ptr.body = fmt.Sprintf(`%s %s SET %s`, verb, query_ptr.tableName, set_body)
	return query_ptr
}

func (query_ptr *Query) QueryStringUpdate() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}
	query += ";"
	return
}
