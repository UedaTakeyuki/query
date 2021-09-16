# query
simple go query builder with just minimum features for me. Pull / Feature requests are welcome as your preferences!

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

columns := []string{"name", "price"}      // column names for select phrase
results := []interface{}{&name, &price} // address of valiables to get value

// SELECT name, price FROM tests WHERE ID = 1;
query = querybuilder.Select(columns).Where(qb.Equal("ID", 1)).QueryString() 
db.SQLiteHandle.QueryRow(query, results...); // The values from the query are set to the variable name and price

/**************************************
* update
***************************************/
// Json Function can be handled
json_func := qb.JsonFunction{Body: fmt.Sprintf(`json_insert(attr, "$.%s", "%s", "$.%s", "%s")`, "kero", "kerokero", "keroyon", "bahahai")}
// name & value pair for update
params = []qb.Param{
	{Name: "name", Value: "frog"},
	{Name: "attr", Value: json_func},
}
// UPDATE tests SET name="frog", attr=json_insert(attr, "$.kero", "kerokero", "$.keroyon", "bahahai") WHERE ID = 1;
query = querybuilder.Update(params).Where(qb.Equal("ID", 1)).QueryString()
db.SQLiteHandle.Exec(query);

/**************************************
* replace into
***************************************/
params := []qb.Param{{Name: "ID",    Value: 1},
                     {Name: "name",  Value: "super cat"},
		                 {Name: "price", Value: 10000}}
		                 {Name: "attr",  Value: "{}"}}
query = querybuilder.ReplaceInto(params).QueryString()
db.SQLiteHandle.Exec(query)
```

## Features
### SQL features
- [x] select
- [x] where
- [x] replace into
- [x] update

### SQL extensions
- [x] Json Function Handling
