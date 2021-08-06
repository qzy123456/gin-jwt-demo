package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"strings"
)

func (m *Middleware) SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "") {
					DebugStack += v + ""
				}
				m.Service.RequestLogger.WithFields(logrus.Fields{
					"ip":            c.ClientIP(), //请求ip
					"method":        c.Request.Method,
					"path":          c.Request.URL.Path,
					"agent":         c.Request.UserAgent(),
					"body":          c.Request.Body,
					"error_message": err,
					"error":         DebugStack,
					"status":        c.Writer.Status(),
				}).Error("Error")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
