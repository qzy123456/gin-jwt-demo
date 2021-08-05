package middleware

import (
	"bytes"
	"fmt"
	"jwtDemo/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (m *Middleware) OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			body   []byte
		)
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		startTime := time.Now()
		//处理请求
		c.Next()
		// 执行时间
		latency := time.Now().Sub(startTime)
		//存储到数据库，忽略error
		if err:=m.Service.SaveOperation(model.Operation{
			Ip:c.ClientIP(),
			Method:c.Request.Method,
			Path :c.Request.URL.Path,
			Body : string(body),
			Response: writer.body.String(),
			//CreateTime :time.Now().Unix(),
		});err!=nil{
			fmt.Println(err)
		}
		//不在不需要打日志的列表内
		if _, ok := m.NoLoginAction[c.Request.URL.Path]; !ok {
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

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
