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

func NotEqual(name string, val interface{}) string {
	return fmt.Sprintf(`%s != %s`, name, ToLiteralValue(val))
}

func GreaterThan(name string, val interface{}) string {
	return fmt.Sprintf(`%s > %s`, name, ToLiteralValue(val))
}

func GreaterEqual(name string, val interface{}) string {
	return fmt.Sprintf(`%s >= %s`, name, ToLiteralValue(val))
}

func LessThan(name string, val interface{}) string {
	return fmt.Sprintf(`%s < %s`, name, ToLiteralValue(val))
}

func LessEqual(name string, val interface{}) string {
	return fmt.Sprintf(`%s <= %s`, name, ToLiteralValue(val))
}

func And(lhs string, rhs string) string {
	return fmt.Sprintf(`(%s AND %s)`, lhs, rhs)
}

func Or(lhs string, rhs string) string {
	return fmt.Sprintf(`(%s OR %s)`, lhs, rhs)
}

func Not(rhs string) string {
	return fmt.Sprintf(`(NOT %s)`, rhs)
}
