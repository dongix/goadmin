package connection

import (
	"config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string

func init() {
	conf := new(config.MysqlConfig)
	error := config.GetConfig(conf)
	if error != nil {
		log.Fatal("load mysql config error!")
		panic(error)
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=90s&collation=utf8mb4_unicode_ci", conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
}

// GetDB 获取连接
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	return db, err
}
