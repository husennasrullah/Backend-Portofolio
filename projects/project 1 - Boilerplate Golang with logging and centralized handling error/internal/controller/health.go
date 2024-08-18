package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/config"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/constants"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

var (
	StartTime = time.Now()
)

type HealthResponse struct {
	Status         string `json:"status"`
	Version        string `json:"version"`
	ServerTime     string `json:"server_time"`
	Uptime         string `json:"uptime"`
	DatabaseStatus string `json:"database_status"`
	RedisStatus    string `json:"redis_status"`
}

func HealthCheck(db *sqlx.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check database connection
		var dbStatus string
		if err := db.Ping(); err != nil {
			dbStatus = constants.STATUS_DOWN
		} else {
			dbStatus = constants.STATUS_UP
		}

		// Calculate uptime
		uptime := time.Since(StartTime).String()

		// Determine overall status
		var overallStatus string
		if dbStatus == constants.STATUS_UP {
			overallStatus = constants.STATUS_UP
		} else {
			overallStatus = constants.STATUS_DOWN
		}

		// Create health response
		response := HealthResponse{
			Status:         overallStatus,
			Version:        cfg.AppVersion,
			ServerTime:     time.Now().Format(time.RFC3339),
			Uptime:         uptime,
			DatabaseStatus: dbStatus,
			RedisStatus:    constants.STATUS_UNKNOWN,
		}
		c.JSON(http.StatusOK, response)
	}
}
