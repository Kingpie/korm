package session

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"korm/dialect"
	"korm/log"
	"testing"
)

func TestSession_Exec(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@/user?charset=utf8")
	if err != nil {
		t.Errorf("open mysql failed, err:%v\n", err)
		return
	}
	defer func() { _ = db.Close() }()

	dealect, _ := dialect.GetDialect("mysql")
	s := New(db, dealect)
	s = s.Raw("select * from score")
	result, err := s.Exec()
	if err != nil {
		t.Errorf("err:%s", err)
		return
	}

	log.Info(result.RowsAffected())
}
