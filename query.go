package query

import (
	"fmt"
)

type Query struct {
	// private
	tableName string
	body      string
	where     string
	err_str   string
}

func (query_ptr *Query) SetTableName(tableName string) *Query {
	query_ptr.tableName = tableName
	return query_ptr
}

func (query_ptr *Query) Select(columns []string) *Query {
	// confirm tableName is not ""
	if query_ptr.tableName == "" {
		query_ptr.err_str += "Need to set tableName before setting verb; "
		return query_ptr
	}

	verb := "SELECT"
	params := ""
	for _, column := range columns {
		if column != "" {
			if params == "" {
				params += fmt.Sprintf(`%s`, column)
			} else {
				params += fmt.Sprintf(` ,%s`, column)
			}
		}
	}
	if params == "" {
		params = "* "
	}
	query_ptr.body = fmt.Sprintf(`%s %s FROM %s`, verb, params, query_ptr.tableName)
	return query_ptr
}

func (query_ptr *Query) GetQuery() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}
	query += ";"
	return
}

func toStr(val interface{}) string {
	switch val := val.(type) {
	case int:
		return fmt.Sprintf(`%d`, val)
	case string:
		return val
	}
	return ""
}
