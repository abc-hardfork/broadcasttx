package routers

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/qshuai/broadcasttx/api"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

// should initial before the first request coming.
var WhiterList []string

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(inAllowIP())
	r.Use(LogContext())

	root := r.Group("service")

	b := root.Group("/broadcast")
	{
		b.POST("/abc", api.BroadcastAbcTx)
		b.POST("/sv", api.BroadcastSvTx)
	}

	root.GET("/fetchtx/:hash", api.FetchTx)

	return r
}

func inAllowIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		var isAllowed bool
		for _, item := range WhiterList {
			if item == ip {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.AbortWithStatusJSON(200, api.ErrNotAllowedIPResponse)
		}

		c.Next()
	}
}

type ResultWriter struct {
	resp *bytes.Buffer
	gin.ResponseWriter
}

func (rw *ResultWriter) Write(p []byte) (int, error) {
	size, err := rw.resp.Write(p)
	if err != nil {
		return size, err
	}
	return rw.ResponseWriter.Write(p)
}

var DefaultTraceLabel = "traceID"

func LogContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Query(DefaultTraceLabel)
		if traceID == "" {
			uniqueID, _ := uuid.NewV4()
			traceID = uniqueID.String()
		}

		// unique track request identification, use c.MustGet("reqID") to
		// get this unique environment variable
		c.Set(DefaultTraceLabel, traceID)

		rw := &ResultWriter{
			resp:           bytes.NewBuffer(nil),
			ResponseWriter: c.Writer,
		}
		c.Writer = rw

		c.Next()

		logrus.WithFields(logrus.Fields{
			"traceID": traceID,
			"IP":      c.ClientIP(),
			"URI":     c.Request.RequestURI,
			"Params":  c.Request.Form,
		}).Info("request information:")

		logrus.WithFields(logrus.Fields{
			"traceID":  traceID,
			"IP":       c.ClientIP(),
			"URI":      c.Request.RequestURI,
			"Params":   c.Request.Form,
			"Response": rw.resp.String(),
		}).Info("response information:")

	}
}
