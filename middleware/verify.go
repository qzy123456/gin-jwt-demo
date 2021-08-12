package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 参数校验中间件
func (m *Middleware) ParamVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := struct{}{}

		validate := validator.New()
		err := validate.Struct(params)
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": -1,
				"msg":  "nil " + err.(validator.ValidationErrors)[0].Field(),
			})
			c.Abort()
			return
		}
	}
}
