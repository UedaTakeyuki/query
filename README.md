# query
simple go query builder with **Json Function** support. Pull / Feature requests are welcome as your preferences!

## How to use
```
import (
  "fmt"
  qb "github.com/UedaTakeyuki/query"
)

// query builder
var querybuilder qb.Query

/**************************************
* set table name first
***************************************/
querybuilder.SetTableName("tests")

/**************************************
* select
***************************************/
// variables to get value from the query
var name string
var price int

columns := []string{"name", "price"}                                           // column names for select phrase
results := []interface{}{&name, &price}                                        // address of valiables to get value

// SELECT name, price FROM tests WHERE ID = 1;
query = querybuilder.Select(columns).Where(qb.Equal("ID", 1)).QueryString() 
db.SQLiteHandle.QueryRow(query, results...);                                   // set values to name and price

/**************************************
* update
***************************************/
//
// Update ordinary column as "name" and Json columng as "attr" by Params struct with 2 elements for "name" and "attr"
//
json_func := qb.JsonFunction{                                                  // Json Functions is supported
               Body:
               fmt.Sprintf(`json_insert(attr, "$.%s", "%s", "$.%s", "%s")`,    // In created SQL string, this json function string is treated as special
			   "kero",                                             // Fx: don't be quoted, even ordinaly string shoud be quoted.
			   "kerokero",
			   "keroyon",
			   "bahahai")
}
			     
// name & value pair for update
params = []qb.Param{
	{Name: "name", Value: "frog"},
	{Name: "attr", Value: json_func},
}
// UPDATE tests SET name="frog", attr=json_insert(attr, "$.kero", "kerokero", "$.keroyon", "bahahai") WHERE ID = 1;
query = querybuilder.Update(params).Where(qb.Equal("ID", 1)).QueryString()
db.SQLiteHandle.Exec(query);

//
// Update Json object value by new json value
//

newJsonStr := '{"d":{"e":100,"f":"ケロケロ"}}'

// for json_set() function
json_func := query.JsonFunction{
	// STR2JSON_FUNC is abstruct keyword, be adjusted to each DBMS automaticallhy
	//   fx: "json" for sqlite 
	//       The json(x) function verifies that its argument X is a valid JSON string and returns a minified version of that JSON string.
	//       "json_compact" for mariadb
	//       The json_compact() function Removes all unnecessary spaces so the json document is as short as possible.
	
	Body: fmt.Sprintf(`json_set(Attr, "$.runtime_settings", STR2JSON_FUNC('%s'))`, newJsonStr)
}

// just for `Attr=json_set()`
params := []query.Param{
	{Name: "Attr", Value: json_func},
}

// UPDATE device SET Attr=json_set(Attr, "$.runtime_settings", STR2JSON_FUNC('{"d":{"e":100,"f":"ケロケロ"}}')) WHERE ID = 'vogQLP';
query = querybuilder.Update(params).Where(query.Equal("ID", 'vogQLP')).QueryString()

/**************************************
* replace into
***************************************/
params := []qb.Param{{Name: "ID",    Value: 1},
                     {Name: "name",  Value: "super cat"},
		                 {Name: "price", Value: 10000}}
		                 {Name: "attr",  Value: "{}"}}
query = querybuilder.ReplaceInto(params).QueryString()
db.SQLiteHandle.Exec(query)

/**************************************
* delete
***************************************/
// DELETE FROM tests WHERE ID = 'kero';
query = querybuilder.Delete().Where(query.Equal("ID", "kero")).QueryString()
db.SQLiteHandle.Exec(query)
```

## Features
### JSON Function support
The **JsonFunction** type is provided to avoid making wrong SQL phrase by which json functions are missused as ordinaly string or other expr.

```
json_func := qb.JsonFunction{                                                  // Json Functions is supported
               Body:
               fmt.Sprintf(`json_insert(attr, "$.%s", "%s", "$.%s", "%s")`,    // In created SQL string, this json function string is treated as special
			   "kero",                                             // Fx: don't be quoted, even ordinaly string shoud be quoted.
			   "kerokero",
			   "keroyon",
			   "bahahai")
}
			     
// name & value pair for update
params = []qb.Param{
	{Name: "name", Value: "frog"},
	{Name: "attr", Value: json_func},
}
// UPDATE tests SET name="frog", attr=json_insert(attr, "$.kero", "kerokero", "$.keroyon", "bahahai") WHERE ID = 1;
query = querybuilder.Update(params).Where(qb.Equal("ID", 1)).QueryString()
```

### ToLiteralValue: Quote a literal value appropriately
The function [ToLiteralValue](https://github.com/UedaTakeyuki/query/blob/main/query.go#L68) quotes a literalvalue in accordance with their type.

```
import (
	qb "github.com/UedaTakeyuki/query"
)

func (data *V1Data) dbUpdate(id interface{}, jsonPath string, value interface{}) (queryStr string) {

	valueStrQuotedApropriately := qb.ToLiteralValue(value)

	json_func := qb.JsonFunction{
		Body: fmt.Sprintf(`json_set(Attr, "%s", %s)`,
			jsonPath,
			valueStrQuotedApropriately),
	}

	params := []qb.Param{
		{Name: "Attr", Value: json_func},
	}

	queryStr = queryBuilder.SetTableName("product").Update(params).Where(query.Equal("ID", id)).QueryString()
	return
}

func showQueries(detail map[string]interface{}){
	name := "cake"
	price := 1.5
	
	log.Println(dbUpdate(1, "$.name", "cake")) // Update a Json column "Attr.name" by string "cake"
	log.Println(dbUpdate(1, "$.price", 1.5))   // Update a Json column "Attr.name" by string "cake"
	
	jsonStr, err := json.Marshal(detail)
	
	// wrap json_func struct for indicate it as "json funcsiont" explicitly
	json_func := qb.JsonFunction{
		// STR2JSON_FUNC is abstruct keyword, be adjusted to each DBMS automaticallhy
		// fx: "json" for sqlite, "json_compact" for mariadb
		Body: fmt.Sprintf(`STR2JSON_FUNC('%s')`, jsonStr),
	}
	
	log.Println(dbUpdate(1. "$.detail", json_func)) // Update a Json column "Attr.name" by a json data created from jsonStr.
	
}
```
### STR2 macro
```STR2xxx``` macro is expanded to appropriate stuff depending　on the DBMS with ```Str2SQLite()``` or ```Str2Mariadb()``` expansion function as follwos:

#### STR2JSON_FUNC() macro
This macro is expanded to a function that tells dbms to treat the argument string as json.
- ```Str2SQLite()``` expand this as ```json()```.
- ```Str2Mariadb()``` expand this as ```json_compact```.

#### STR2PF macro
This macro is expanded to the placeholder string of prepared statements on the DBMS.
- ```Str2SQLite()``` expand this as ```?```.
- ```Str2Mariadb()``` expand this as ```?```.
Note that placeholder string happens to be the same for both SQLite and Mariadb, but will likely be replaced with a different string in the future when postgres and oracle are supported.

#### STR2PF_PATH macro
This macro is expanded to the placeholder string for the json_path.
- ```Str2SQLite()``` expand this as ```'' || ?```.
- ```Str2Mariadb()``` expand this as ```CONCAT('', ?)```.
You may think why so complecated. For more detail, refer the [Sudip Raval](https://medium.com/@rsudip90)'s [blog](https://medium.com/aubergine-solutions/working-with-mysql-json-data-type-with-prepared-statements-using-it-in-go-and-resolving-the-15ef14974c48) and [his article of stackoverflow](https://stackoverflow.com/questions/41436245/how-to-use-a-prepared-statement-for-json-data-types-for-mysql-in-java)

### Expediently feature support
Basically, supported features are selected to meet my necesity for my projects :-)  
Although, feature request are welcome!


## Status
### SQL features
- [x] select
- [x] where
- [x] replace into
- [x] insert into
- [x] update
- [x] drop database
- [x] delete

### SQL extensions
- [x] Json Function Handling
