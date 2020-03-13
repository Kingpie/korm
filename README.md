# korm
Simple ORM For Go  

quick start:
```go
type Score struct {
	Id   uint64 `korm:"PRIMARY KEY"`
	Name string `korm:"not null"`
	Age  int8   `korm:"not null default 1"`
}

func main(){
	engine, _ := NewEngine("mysql", "root:123456@/User?charset=utf8")
	defer engine.Close()
	s := engine.NewSession().Model(&Score{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_,_ = s.Insert(&Score{1, "tom", 3})
	score := &Score{}
	_ = s.First(score)
	
	scoreList := []Score{}
	_ = s.Find(&scoreList)
}
```