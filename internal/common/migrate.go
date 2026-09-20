package common

import (
	"embed"

	"github.com/pressly/goose/v3"
)

func InitMigration(fs embed.FS) {
	goose.SetBaseFS(fs)
	_ = goose.SetDialect("sqlite3")
}
