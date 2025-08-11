package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"korm/dialect"
	"korm/log"
	"testing"
)

type account struct {
	ID       int64 `korm:"PRIMARY KEY"`
	Password string
}

func (account *account) BeforeInsert(s *Session) error {
	log.Info("before insert", account)
	account.ID += 100000
	return nil
}

func (account *account) AfterQuery(s *Session) error {
	log.Info("after query", account)
	account.Password = "********"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@/user?charset=utf8")
	dia, _ := dialect.GetDialect("mysql")
	defer db.Close()

	s := New(db, dia).Model(&account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&account{1, "123"}, &account{2, "456"}, &account{3, "666"})

	req := []account{}
	s.Find(&req)
	t.Logf("%+v", req)
}
