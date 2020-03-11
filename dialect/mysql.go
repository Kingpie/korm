package dialect

import (
	"fmt"
	"reflect"
)

type mysql struct{}

var _ Dialect = (*mysql)(nil)

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(val reflect.Value) string {
	switch val.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int16, reflect.Int32:
		return "int"
	case reflect.Int8:
		return "tinyint"
	case reflect.Uint8:
		return "tinyint unsigned"
	case reflect.Uint, reflect.Uint16, reflect.Uint32:
		return "int unsigned"
	case reflect.Int64:
		return "bigint"
	case reflect.Uint64:
		return "bigint unsigned"
	case reflect.Float32, reflect.Float64:
		return "double"
	case reflect.String:
		return "text"
	}

	panic(fmt.Sprintf("invalid sql type %s %s", val.Type().Name(), val.Kind()))
}

func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT table_name FROM information_schema.TABLES WHERE table_name = ? ;", args
}
