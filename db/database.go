package db

import (
	"database/sql"
	"embed"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var dbMigrations embed.FS

var DbConnection *sql.DB

func InitDB(connectionString string) error {
	var err error
	DbConnection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	if err = DbConnection.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}

func DBMigrate() error {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	n, err := migrate.Exec(DbConnection, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
