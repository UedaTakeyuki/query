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

// json() function string should be separated from ordinaly string
// because it shouldn't quoted in the SQL string.
// This difference is appear in the ToLiteralValue() function.
type JsonFunction struct {
	Body string
}

const (
	Select Verb = iota
	ReplaceInto
	InsertInto
	Update
	DropTable
	Delete
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
	case InsertInto:
		return query_ptr.QueryStringInsertInto()
	case Update:
		return query_ptr.QueryStringUpdate()
	case DropTable:
		return query_ptr.QueryStringDropTable()
	case Delete:
		return query_ptr.QueryStringDelete()
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
	case int64:
		return fmt.Sprintf(`%d`, val)
	case string:
		return fmt.Sprintf(`'%s'`, val)
	case JsonFunction:
		return fmt.Sprintf(`%s`, val.Body)
	}
	return ""
}
