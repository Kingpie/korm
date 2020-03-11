package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

//封装多种数据库
type Dialect interface {
	DataTypeOf(val reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
