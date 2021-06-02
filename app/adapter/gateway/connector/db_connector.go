package connector

import "database/sql"

type DBConnector interface {
	GetDB() *sql.DB
}
