package schema

import (
	"github.com/Kingpie/korm/dialect"
	"testing"
)

type Score struct {
	Id   string `korm:"PRIMARY KEY"`
	Name string
	Age  int
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParse(t *testing.T) {
	table := Parse(&Score{}, TestDial)
	t.Logf("%+v", *table)
}
