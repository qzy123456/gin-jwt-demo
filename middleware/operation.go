package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jwtDemo/model"
	"jwtDemo/utils"
	"net/http"
	"time"
)

func (m *Middleware) OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body []byte
		)
		writer := ResponseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		startTime := time.Now()
		// 执行时间
		latency := time.Now().Sub(startTime)
		var err error
		body, err = ioutil.ReadAll(c.Request.Body)
		if err == nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		//不在不需要打日志的列表内
		if _, ok := m.NoLoginAction[c.Request.URL.Path]; !ok {
			//存储到数据库，忽略error
			if c.Request.Method == http.MethodPost {
				m.Service.SaveOperation(model.Operation{
					Ip:         c.ClientIP(),
					Method:     c.Request.Method,
					Path:       c.Request.URL.Path,
					Body:       string(body),
					Response:   writer.body.String(),
					CreateTime: utils.GetYmds(),
				})
			}

			// 日志格式
			m.Service.RequestLogger.WithFields(logrus.Fields{
				"ip":            c.ClientIP(), //请求ip
				"method":        c.Request.Method,
				"path":          c.Request.URL.Path,
				"agent":         c.Request.UserAgent(),
				"body":          string(body),
				"error_message": c.Errors.ByType(gin.ErrorTypePrivate).String(),
				"status":        c.Writer.Status(),
				"latency":       latency,
				"response":      writer.body.String(),
			}).Info("OperationRecord")
		}
	}
}

type ResponseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r ResponseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
