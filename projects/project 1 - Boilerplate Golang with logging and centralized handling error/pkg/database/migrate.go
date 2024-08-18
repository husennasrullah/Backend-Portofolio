package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/husennasrullah/Backend-Portofolio/project-1/config"
	_ "github.com/lib/pq"
	"log"
)

func MigrationUp(cfg *config.Config) error {
	// Format string koneksi database
	dbAddress := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("postgres", dbAddress)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	defer db.Close()

	//create driver postgres
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}

	// Path to migration directory
	migrationsPath := "file://utils/database/schema_migration"

	// Membuat instance migrasi
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, cfg.DBName, driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// migrate up
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}
