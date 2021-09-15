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
	verb      Verb
}

type Verb int

const (
	Select Verb = iota
	ReplaceInto
)

func (query_ptr *Query) SetTableName(tableName string) *Query {
	query_ptr.tableName = tableName
	return query_ptr
}

func (query_ptr *Query) QueryString() (query string) {
	switch query_ptr.verb {
	case Select:
		return query_ptr.QueryStringSelect()
	case ReplaceInto:
		return query_ptr.QueryStringReplaceInto()
	default:
		query_ptr.err_str += "invalidate verb; "
		return ""
	}
	return
}

/*
	Make literal-value from go variable

	.(int)    val
	.(string) "val"
*/
func ToLiteralValue(val interface{}) string {
	switch val := val.(type) {
	case int:
		return fmt.Sprintf(`%d`, val)
	case string:
		return fmt.Sprintf(`"%s"`, val)
	}
	return ""
}
