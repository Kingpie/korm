package session

import (
	"database/sql"
	"github.com/Kingpie/korm/dialect"
	"github.com/Kingpie/korm/log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSession_Exec(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@/User?charset=utf8")
	defer func() { _ = db.Close() }()

	dealect, _ := dialect.GetDialect("mysql")
	s := New(db, dealect)
	s = s.Raw("select * from test")
	result, err := s.Exec()
	if err != nil {
		t.Errorf("err:%s", err)
		return
	}

	log.Info(result.RowsAffected())
}
