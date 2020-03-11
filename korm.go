package korm

import (
	"database/sql"
	"github.com/Kingpie/korm/dialect"
	"github.com/Kingpie/korm/log"
	"github.com/Kingpie/korm/session"
)

//与用户交互
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("dialect %s not found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
	log.Info("Connect db success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("close db failed")
	}

	log.Info("close db success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
