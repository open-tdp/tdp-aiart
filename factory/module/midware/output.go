package midware

import (
	"github.com/gin-gonic/gin"
)

func OutputHandle() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

		// 输出错误信息

		if err, exists := c.Get("Error"); exists {
			c.AbortWithStatusJSON(exitCode(c, 400), newErrorMessage(err))
			return
		}

		// 输出请求结果

		msg := c.GetString("Message")

		if res, exists := c.Get("Payload"); msg != "" || exists {
			c.AbortWithStatusJSON(exitCode(c, 200), newPayload(res, msg))
			return
		}

		// 输出HTML内容

		if htm := c.GetString("HTML"); htm != "" {
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, htm)
			c.Abort()
			return
		}

		// 捕获异常返回

		c.AbortWithStatusJSON(500, newErrorMessage("内部错误"))

	}

}
