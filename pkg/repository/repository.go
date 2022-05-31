package repository

import (
	"database/sql"
	"rs_bot/pkg/config"
)

type Repository interface {
	DbConnection(cfg config.ConfigEnv) (*sql.DB, error)
	dsn(dbName string, cfg config.ConfigEnv) string
}
