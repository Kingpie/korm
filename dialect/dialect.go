package dialect

import "reflect"

var dialectMap = map[string]Dialect{}

// 封装多种数据库
type Dialect interface {
	DataTypeOf(val reflect.Value) string                    //用于将 Go 语言的类型转换为该数据库的数据类型。
	TableExistSQL(tableName string) (string, []interface{}) //返回某个表是否存在的 SQL 语句
}

func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
