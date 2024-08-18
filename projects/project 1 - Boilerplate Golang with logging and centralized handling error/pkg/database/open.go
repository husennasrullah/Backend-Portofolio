package database

import (
	"fmt"
	"github.com/husennasrullah/Backend-Portofolio/project-1/config"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/database/drivers"
	"github.com/jmoiron/sqlx"
	"time"
)

const DRIVER_POSTGRES = "postgres"

func SetupPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
	dbHost := cfg.DBHost
	dbUser := cfg.DBUser
	dbPass := cfg.DBPassword
	dbName := cfg.DBName
	dbPort := cfg.DBPort
	sslMode := cfg.SSLMode
	tz := cfg.TZ

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslMode, tz,
	)

	// Setup sqlx config of postgreSQL
	configDB := drivers.SQLXConfig{
		DriverName:     DRIVER_POSTGRES,
		DataSourceName: dsn,
		MaxOpenConns:   100,
		MaxIdleConns:   10,
		MaxLifetime:    15 * time.Minute,
	}

	// Initialize postgreSQL connection with sqlx
	conn, err := configDB.InitializeSQLXDatabase()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
