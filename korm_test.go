package korm

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"korm/session"
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

func OpenDB(t *testing.T) *Engine {
	engine, _ := NewEngine("mysql", "root:123456@/User?charset=utf8")
	return engine
}

func TestEngine_TransactionRollback(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&Score{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (i interface{}, e error) {
		_ = s.Model(&Score{}).CreateTable()
		_, e = s.Insert(&Score{1, "fff", 3})
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func TestEngine_TransactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&Score{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (i interface{}, e error) {
		_ = s.Model(&Score{}).CreateTable()
		_, e = s.Insert(&Score{1, "fff", 3})
		return
	})

	score := &Score{}
	_ = s.First(score)
	if err != nil {
		t.Fatal("fail to commit")
	}
}
