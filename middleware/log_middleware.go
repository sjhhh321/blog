package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}
func LogMiddleware(c *gin.Context) {
	// 请求中间件
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	fmt.Println("body: ", string(byteData))
	// 这是因为c.Request.Body被使用后就会销毁，不能再绑定，所以需要重新赋值创建。
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	res := &ResponseWriter{
		ResponseWriter: c.Writer,
	}
	c.Writer = res
	c.Next()
	// 响应中间件
	fmt.Println("response: ", string(res.Body))
}
