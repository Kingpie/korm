package schema

import (
	"github.com/Kingpie/korm/dialect"
	"go/ast"
	"reflect"
)

//列
type Field struct {
	Name string //字段名
	Type string //字段类型
	Tag  string //约束条件
}

//表
type Schema struct {
	Model      interface{}       //映射对象
	Name       string            //表名
	Fields     []*Field          //字段
	FieldNames []string          //字段名
	fieldMap   map[string]*Field //字段名与字段映射
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

//将结构体解析为Schema
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}

			if v, ok := p.Tag.Lookup("korm"); ok {
				field.Tag = v
			}

			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}

	return schema
}
