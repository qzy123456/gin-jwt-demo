package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
)

func (m *Middleware) SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body  []byte
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "") {
					DebugStack += v + ""
				}
				body, err = ioutil.ReadAll(c.Request.Body)
				if err == nil {
					c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				}
				m.Service.RequestLogger.WithFields(logrus.Fields{
					"ip":            c.ClientIP(), //请求ip
					"method":        c.Request.Method,
					"path":          c.Request.URL.Path,
					"agent":         c.Request.UserAgent(),
					"body":          string(body),
					"error_message": err,
					"error":         DebugStack,
				}).Error("Error")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
