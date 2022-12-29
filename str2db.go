package query

import "strings"

func Str2SQLite(queryString string) (replacedString string) {
	replacedString = queryString
	if strings.Index(replacedString, "STR2JSON_FUNC") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2JSON_FUNC", "json", -1) // -1: no limit
	}
	if strings.Index(replacedString, "STR2PF_PATH") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2PF_PATH", "'' || ?", -1) // -1: no limit
	}
	if strings.Index(replacedString, "STR2PF") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2PF", "?", -1) // -1: no limit
	}
	return
}

func Str2Mariadb(queryString string) (replacedString string) {
	replacedString = queryString
	if strings.Index(replacedString, "STR2JSON_FUNC") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2JSON_FUNC", "json_compact", -1) // -1: no limit
	}
	if strings.Index(replacedString, "STR2PF_PATH") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2PF_PATH", "CONCAT('', ?)", -1) // -1: no limit
	}
	if strings.Index(replacedString, "STR2PF") != -1 { // queryStr inclues the string
		replacedString = strings.Replace(replacedString, "STR2PF", "?", -1) // -1: no limit
	}
	return
}
