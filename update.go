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

	// set verb as "ReplaceInto"
	query_ptr.verb = Update
	verb := "UPDATE"
	set_body := ""

	// confirm params
	if params != nil {

		if len(params) == 0 {
			query_ptr.err_str += "params of ReplaceInto should not blank array; "
			return query_ptr
		}

		// make body
		for _, param := range params {
			if set_body == "" {
				set_body += fmt.Sprintf(`%s=%s`, param.Name, ToLiteralValue(param.Value))
			} else {
				set_body += fmt.Sprintf(`,%s=%s`, param.Name, ToLiteralValue(param.Value))
			}
		}

		query_ptr.body = fmt.Sprintf(`%s %s SET %s`, verb, query_ptr.tableName, set_body)
		return query_ptr
	} else {
		// param not set, so SET clause can be set later
		query_ptr.body = fmt.Sprintf(`%s %s`, verb, query_ptr.tableName)
		return query_ptr
	}
}

func (query_ptr *Query) QueryStringUpdate() (query string) {
	query = fmt.Sprintf(`%s`, query_ptr.body)
	if query_ptr.where != "" {
		query += fmt.Sprintf(` %s`, query_ptr.where)
	}
	query += ";"
	return
}

func (query_ptr *Query) Set(params []Param) *Query {
	set_body := ""

	// make body
	for _, param := range params {
		if set_body == "" {
			set_body += fmt.Sprintf(`%s=%s`, param.Name, ToLiteralValue(param.Value))
		} else {
			set_body += fmt.Sprintf(`,%s=%s`, param.Name, ToLiteralValue(param.Value))
		}
	}

	query_ptr.body = query_ptr.body + fmt.Sprintf(` SET %s`, set_body)
	return query_ptr
}

type JsonPathAndValue Pair struct {
	Path  string      // expression of Json Path
	Value interface{}
}

func (query_ptr *Query) SetJson_Set(path string, jsonParams []JsonPathAndValue) *Query {
	set_body := ""

	// make body
	if set_body == "" {
		set_body += fmt.Sprintf(`%s=%s`, jsonParams.Path, ToLiteralValue(jsonParams.Value))
	} else {
		set_body += fmt.Sprintf(`,%s=%s`, jsonParams.Path, ToLiteralValue(jsonParams.Value))
	}

}