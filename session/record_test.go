package session

import (
	"database/sql"
	"github.com/Kingpie/korm/dialect"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type test struct {
	Name  string
	Score int64
}

var (
	test1 = &test{"hehe", 1}
	test2 = &test{"lala", 2}
	test3 = &test{"haha", 3}
)

func TestSession_Find(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@/User?charset=utf8")
	dia, _ := dialect.GetDialect("mysql")
	defer db.Close()

	s := New(db, dia).Model(&test{})
	err := s.DropTable()
	t.Logf("%s", err)
	_ = s.CreateTable()
	_, _ = s.Insert(test1, test2, test3)

	var users []test
	err = s.Find(&users)
	if err != nil {
		t.Fatal("fail to query")
	}

	t.Logf("%+v", users)
}

func TestSession_Update(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@/User?charset=utf8")
	dia, _ := dialect.GetDialect("mysql")
	defer db.Close()

	s := New(db, dia).Model(&test{})
	_ = s.DropTable()
	//t.Logf("%s", err)
	_ = s.CreateTable()
	_, _ = s.Insert(test1, test2, test3)

	affected, _ := s.Where("Name = ?", "hehe").Update("Score", 111)
	u := &test{}
	_ = s.OrderBy("Score DESC").First(u)

	t.Logf("%d,%+v", affected, u)
}

func TestSession_DeleteAndCount(t *testing.T) {
	db, _ := sql.Open("mysql", "root:123456@/User?charset=utf8")
	dia, _ := dialect.GetDialect("mysql")
	defer db.Close()

	s := New(db, dia).Model(&test{})
	_ = s.DropTable()
	//t.Logf("%s", err)
	_ = s.CreateTable()
	_, _ = s.Insert(test1, test2, test3)

	count, _ := s.Count()
	t.Logf("count=%d", count)
	affected, _ := s.Where("Name=?", "hehe").Delete()
	count, _ = s.Count()
	t.Logf("affected:%d,count=%d", affected, count)
}
