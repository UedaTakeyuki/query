package query

import (
	"fmt"
)

func (query_ptr *Query) Where(cond string) *Query {
	// confirm cond shoud not ""
	if cond == "" {
		query_ptr.err_str += `condition of WHERE should not ""; `
		return query_ptr
	}
	query_ptr.where = fmt.Sprintf(`WHERE %s`, cond)
	return query_ptr
}

/*
  Make Equal condition expr with literal-value from go variable

  val.(type) == int:      name = val
	val.(type) == string    name = "val"
*/
func Equal(name string, val interface{}) string {
	return fmt.Sprintf(`%s = %s`, name, ToLiteralValue(val))
}
