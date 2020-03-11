package korm

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type Score struct {
	Id   uint64 `korm:"PRIMARY KEY"`
	Name string `korm:"not null"`
	Age  int8   `korm:"not null default 1"`
}

func TestNewEngine(t *testing.T) {
	engine, _ := NewEngine("mysql", "root:123456@/User?charset=utf8")
	defer engine.Close()

	s := engine.NewSession().Model(&Score{})
	_ = s.DropTable()
	_ = s.CreateTable()

	if !s.HasTable() {
		t.Fatal("failed to create table")
	}
}
