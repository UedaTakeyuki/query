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

func Like(name interface{}, val interface{}) string {
	return fmt.Sprintf(`%s LIKE %s`, ToLiteralValue(name), ToLiteralValue(val))
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

// WHERE (age BETWEEN 40 AND 55);
func Between(name string, val1 interface{}, val2 interface{}) string {
	return fmt.Sprintf(`(%s BETWEEN %s AND %s)`, name, ToLiteralValue(val1), ToLiteralValue(val2))
}

// WHERE (age IN(28, 38, 48));
func In(name string, vals []interface{}) string {
	res := fmt.Sprintf(`(%s IN(`, name)
	for _, val := range vals {
		res = res + fmt.Sprintf(`%s,`, ToLiteralValue(val))
	}
	res = res[0 : len(res)-1]
	res = res + "))"
	return res
}

// WHERE (c1 IS NULL) ;
func IsNull(name interface{}) string {
	return fmt.Sprintf(`(%s IS NULL)`, ToLiteralValue(name))
}

// WHERE (c1 IS NOT NULL) ;
func IsNotNull(name interface{}) string {
	return fmt.Sprintf(`(%s IS NOT NULL)`, ToLiteralValue(name))
}
