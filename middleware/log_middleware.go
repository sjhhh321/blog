package middleware

import (
	"blogx_server/service/log_service"
	"github.com/gin-gonic/gin"
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
	log := log_service.NewActionLogByGin(c)
	log.SetRequest(c)
	c.Set("log", log)
	res := &ResponseWriter{
		ResponseWriter: c.Writer,
	}
	c.Writer = res
	c.Next()
	// 响应中间件
	log.SetResponse(res.Body)
	log.Save()
}
