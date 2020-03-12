package clause

import "testing"

func TestClause_Set(t *testing.T) {
	var clause Clause
	clause.Set(SELECT, "Test", []string{"hehe"})
	clause.Set(WHERE, "Name = ?", "lala")
	clause.Set(ORDERBY, "SCORE DESC")
	clause.Set(LIMIT, 1)
	sql, _ := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Logf("sql=%s", sql)
}
