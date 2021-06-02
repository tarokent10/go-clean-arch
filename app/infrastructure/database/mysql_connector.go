package database

import (
	"database/sql"
	"fmt"
	"picture-go-app/adapter/gateway/connector"
	"picture-go-app/infrastructure/config"
	"picture-go-app/infrastructure/log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Type = "mysql"
)

type MySQLConnector struct {
	db *sql.DB
}

func (con *MySQLConnector) GetDB() *sql.DB {
	return con.db
}

func Open() (connector.DBConnector, error) {
	c := config.Get()

	sqlinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DBUSER, c.DBPASSWORD, c.DBHOST, c.DBPORT, c.DATABASE)
	log.Debug(sqlinfo)
	db, err := sql.Open(Type, sqlinfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLConnector{db: db}, nil
}
