package api

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/husennasrullah/Backend-Portofolio/project-1/pkg/logger"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

type (
	bodyLogWriter struct {
		gin.ResponseWriter
		body *bytes.Buffer
	}
)

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.SetFormatter(&logger.CustomFormatter{})
		logrus.SetLevel(logrus.InfoLevel)

		startTime := time.Now()

		// Proses request
		c.Next()
		reqBody, _ := c.Get("RequestBody")
		resBody, _ := c.Get("ResponseBody")

		latency := time.Since(startTime)

		entry := logrus.WithFields(logrus.Fields{
			"status":        c.Writer.Status(),
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"latency":       latency.Nanoseconds(),
			"error":         c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"user_agent":    c.Request.UserAgent(),
			"request_body":  reqBody,
			"response_body": resBody,
		})

		if len(c.Errors) > 0 {
			entry.Error("HTTP request log")
		} else {
			entry.Info("HTTP request log")
		}

	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request
		// Create a custom response writer
		responseWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = responseWriter

		// Log the request details
		body := "-"
		bodyBytes, err := io.ReadAll(req.Body)
		if err == nil && string(bodyBytes) != "" {
			body = string(bodyBytes)
		}

		c.Next()
		errMesg := "-"
		if len(c.Errors) > 0 {
			errMesg = c.Errors.String()
		}
		c.Set("ServiceError", errMesg)
		c.Set("ResponseWriter", responseWriter.Status())
		c.Set("RequestBody", body)
		c.Set("ResponseBody", responseWriter.body.String())
	}
}
