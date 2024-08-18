package drivers

import (
	"fmt"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/constants"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

type SQLXConfig struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

func (config *SQLXConfig) InitializeSQLXDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	logger.Info(fmt.Sprintf("Setting maximum number of open connections to %d", config.MaxOpenConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetMaxOpenConns(config.MaxOpenConns)

	// set maximum number of idle connections in the pool
	logger.Info(fmt.Sprintf("Setting max number of idle connections to %d", config.MaxIdleConns), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetMaxIdleConns(config.MaxIdleConns)

	// set maximum time to wait for new connection
	logger.Info(fmt.Sprintf("Setting max lifetime for a connection to %s", config.MaxLifetime), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryDatabase})
	db.SetConnMaxLifetime(config.MaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error ping database: %v", err)
	}

	return db, nil
}
