package korm

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestNewEngine(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:123456@/User?charset=utf8")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	t.Logf("Exec success %d affected", count)
}
