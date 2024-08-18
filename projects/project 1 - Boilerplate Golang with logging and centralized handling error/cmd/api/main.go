package main

import (
	"context"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/config"
	"github.com/husennasrullah/Backend-Portofolio/project-1/docs"
	"github.com/husennasrullah/Backend-Portofolio/project-1/internal/server/api"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/constants"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/database"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	//load config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("failed to load configuration", logrus.Fields{
			"error message": err,
		})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	//open connection sqlx
	db, err := database.SetupPostgresConnection(cfg)
	if err != nil {
		logger.Fatal("err initialize db..", logrus.Fields{
			"error message": err,
		})
	}

	//migrate Database
	if cfg.Mode == constants.LOCAL_MODE {
		err = database.MigrationUp(cfg)
		if err != nil {
			logger.Fatal("error to migrate sql : ", logrus.Fields{
				"error message": err,
			})
		}
	}
	//declare wait group for all process
	waitGroup, ctx := errgroup.WithContext(ctx)

	//start http server
	setSwaggerInfo()
	runHttpServer(ctx, waitGroup, cfg, db)

	err = waitGroup.Wait()
	if err != nil {
		logger.Fatal("error from wait group", nil)
	}
}

func runHttpServer(ctx context.Context, waitGroup *errgroup.Group, cfg *config.Config, db *sqlx.DB) {
	ginEngine := gin.New()

	ginEngine.Use(api.LoggingMiddleware())
	ginEngine.Use(api.ErrorMiddleware(), gin.Recovery())

	configCors := cors.DefaultConfig()
	configCors.AddAllowHeaders(constants.HEADER_AUTHORIZATION)
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE", "PATCH"}
	ginEngine.Use(cors.New(configCors))

	api.Register(ginEngine, db, cfg, ctx)
	srv := &http.Server{
		Addr:    ":" + cfg.LocalPort,
		Handler: ginEngine,
	}

	waitGroup.Go(func() error {
		logger.Info("start HTTP server ", logrus.Fields{
			"port": srv.Addr,
		})

		err := srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			logger.Error("HTTP server failed to serve", nil)
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		logger.Info("graceful shutdown HTTP server", nil)
		err := srv.Shutdown(context.Background())
		if err != nil {
			logger.Info("failed to shutdown HTTP server", nil)
			return err
		}
		logger.Info("HTTP server is stopped", nil)
		return nil
	})
}

func setSwaggerInfo() {
	docs.SwaggerInfo.Title = "Golang Boilerplate"
	docs.SwaggerInfo.Description = "Golang Boilerplate"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
}
