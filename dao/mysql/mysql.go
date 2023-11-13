package mysql

import (
	"TieTie/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MYSQLConfig) (err error) {
	dsn := cfg.User + ":" + cfg.Pwd + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.DB + "?parseTime=true&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	//同时打开的连接数，使用中+空闲
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	//最多空闲的连接数，默认为2
	db.SetMaxIdleConns((cfg.MaxIdleConns))

	return
}

func Close() (err error) {
	err = db.Close()
	return
}
