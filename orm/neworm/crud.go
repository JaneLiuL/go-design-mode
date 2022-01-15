package neworm

import "database/sql"

type NewOrmEngine struct {
	DB        *sql.DB
	TableName string
	Exec      []interface{}
	Limit     string
	Where     []string
}

// connect
