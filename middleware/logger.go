package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

type writer struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w writer) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w writer) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogRequest(c)
		logger := &writer{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = logger
		start := time.Now()
		c.Next()
		LogResponse(c, logger, start)
	}
}

func LogRequest(c *gin.Context) {
	path := c.Request.URL.Path
	clientIP := c.ClientIP()
	clientUserAgent := c.Request.UserAgent()
	referer := c.Request.Referer()
	method := c.Request.Method
	requestID := c.Request.Header.Get("x-request-id")

	logrus.WithFields(logrus.Fields{
		"client_ip":  clientIP,
		"method":     method,
		"path":       path,
		"referer":    referer,
		"user_agent": clientUserAgent,
		"request_id": requestID,
	}).Log(logrus.InfoLevel, "request received")
}

func LogResponse(c *gin.Context, writer *writer, start time.Time) {
	dataLength := c.Writer.Size()
	if dataLength < 0 {
		dataLength = 0
	}
	stop := time.Since(start)
	latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
	statusCode := c.Writer.Status()
	body := writer.body.String()
	requestID := c.Request.Header.Get("x-request-id")

	logrus.WithFields(logrus.Fields{
		"status_code": statusCode,
		"data_length": dataLength,
		"body":        body,
		"latency":     latency,
		"request_id":  requestID,
	}).Log(logrus.InfoLevel, "response sent")
}
