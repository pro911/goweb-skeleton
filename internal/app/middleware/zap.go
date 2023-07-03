package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pro911/golang-demo/internal/pkg/biz/logger"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// ZapLogger 接收gin框架默认的日志
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		post := ""

		if c.Request.Method == "POST" {
			// 把request的内容读取出来
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body.Close()
			// 把刚刚读出来的再写进去
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			switch c.ContentType() {
			case "application/json":
				var result map[string]interface{}
				d := jsoniter.NewDecoder(bytes.NewReader(bodyBytes))
				d.UseNumber()
				if err := d.Decode(&result); err == nil {
					bt, _ := jsoniter.Marshal(result)
					post = string(bt)
				}
			default:
				post = string(bodyBytes)
			}
		}

		c.Next()

		cost := time.Since(start)
		logger.ZapInfo(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("post", post),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// ZapRecovery recover项目可能出现的panic，并使用zap记录相关日志
func ZapRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.ZapError(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					logger.ZapError("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.ZapError("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
