package session

import (
	"database/sql"
	"fmt"
	"korm/clause"
	"korm/dialect"
	"korm/log"
	"korm/schema"
	"reflect"
	"strings"
)

// 与db交互
type Session struct {
	db       *sql.DB
	sql      strings.Builder
	sqlVars  []interface{}
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	tx       *sql.Tx
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}

	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}

	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.RefTable().Name)).Exec()
	return err
}

// HasTable 检查数据库中是否存在当前会话关联的表
// 返回值:
//
//	bool - 如果表存在则返回true，否则返回false
func (s *Session) HasTable() bool {
	// 生成检查表是否存在的SQL语句和参数值
	sqlStr, values := s.dialect.TableExistSQL(s.RefTable().Name)

	// 执行查询并获取结果行
	row := s.Raw(sqlStr, values...).QueryRow()

	// 扫描查询结果
	var tmp string
	_ = row.Scan(&tmp)

	// 比较查询结果与表名来判断表是否存在
	return tmp == s.RefTable().Name
}

// Raw appends raw SQL and values to current session
// sql: raw SQL string to be appended
// values: variadic interface{} values to be appended as SQL variables
// returns: pointer to current Session instance for method chaining
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	// Append SQL string and a space to session's SQL buffer
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	// Append provided values to session's SQL variables slice
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 执行sql语句
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}

	return
}

// 查询单条记录
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// 查询多条记录
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
