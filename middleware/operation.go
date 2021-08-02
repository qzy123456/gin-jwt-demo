package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"jwtDemo/model"
)

func (m *Middleware) OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body   []byte
			params model.BaseParams
		)
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				m.Service.RequestLogger.WithFields(logrus.Fields{
					"ip":     c.ClientIP(), //请求ip
					"method": c.Request.Method,
					"path":   c.Request.URL.Path,
					"agent":  c.Request.UserAgent(),
					"error":  err.Error(),
				}).Error("OperationRecord [ioutil.ReadAll] middleware err")
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				err = json.Unmarshal(body, &params)
				if err != nil {
					m.Service.RequestLogger.WithFields(logrus.Fields{
						"body":   string(body),
						"error":  err.Error(),
						"ip":     c.ClientIP(), //请求ip
						"method": c.Request.Method,
						"path":   c.Request.URL.Path,
						"agent":  c.Request.UserAgent(),
					}).Error("OperationRecord [json.Unmarshal] middleware err")
				}
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		startTime := time.Now()
		//处理请求
		c.Set("params", params)
		c.Next()
		// 执行时间
		latency := time.Now().Sub(startTime)

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
		//存储到数据库
		m.Service.SaveOperation(model.Operation{
			Ip:c.ClientIP(),
			Method:c.Request.Method,
			Path :c.Request.URL.Path,
			Body : string(body),
			Response: writer.body.String(),
			CreateTime :time.Now().Unix(),
		})
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
