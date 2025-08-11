package schema

import (
	"go/ast"
	"korm/dialect"
	"reflect"
)

// 列
type Field struct {
	Name string //字段名
	Type string //字段类型
	Tag  string //约束条件
}

// 表
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

// 解析Go结构体并生成对应的数据库模式
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	//通过反射获取传入结构体的类型信息
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), //表名
		fieldMap: make(map[string]*Field),
	}

	//遍历结构体的所有字段
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		//为每个可导出的非匿名字段创建Field对象，记录字段名和对应的数据库类型
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,                                              //字段名
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))), //字段类型
			}

			//处理字段的"korm"标签
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

// RecordValues 将目标结构体中与Schema字段对应的所有字段值提取到一个interface{}切片中
// 参数:
//
//	dest - 指向结构体的指针，包含需要提取的字段值
//
// 返回值:
//
//	[]interface{} - 包含按Schema字段顺序提取的字段值的切片
func (Schema *Schema) RecordValues(dest interface{}) []interface{} {
	// 获取目标结构体的反射值
	destVal := reflect.Indirect(reflect.ValueOf(dest))
	var fieldVals []interface{}

	// 遍历Schema中的所有字段，按字段名提取对应的值
	for _, field := range Schema.Fields {
		fieldVals = append(fieldVals, destVal.FieldByName(field.Name).Interface())
	}

	return fieldVals
}
