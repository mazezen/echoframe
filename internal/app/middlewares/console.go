package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mazezen/echoframe/vm"
	"go.uber.org/zap"
	"runtime"
	"strings"
)

var RequestId = "echo-frame"

var ConsoleBodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: func(c echo.Context) bool {
		if !strings.HasPrefix(c.Path(), "/api/") {
			return true
		}
		return false
	},
	Handler: func(c echo.Context, bytes []byte, bytes2 []byte) {
		if !strings.HasPrefix(c.Path(), "/api/") {
			return
		}
		uid := c.Get(RequestId).(string)
		vm.Box.Log.Info("请求结束", zap.String("请求UID", uid), zap.String(c.Request().RequestURI, string(bytes2)))
	},
}

func FrameLog() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10)
					l := runtime.Stack(stack, false)
					vm.Box.Log.Error("程序崩溃", zap.String("崩溃日志", string(stack[:l])))
				}
			}()
			if !strings.HasPrefix(c.Path(), "/api/") {
				vm.Box.Log.Info("请求开始", zap.Any(c.Request().RequestURI, "网页请求"))
				return handlerFunc(c)
			}
			uid := uuid.New().String()
			c.Set(RequestId, uid)
			err := handlerFunc(c)
			return err
		}
	}
}

var EchoLog *EchoLogger

type EchoLogger struct{}

func (*EchoLogger) Write(p []byte) (n int, err error) {
	vm.Box.Log.Info("ECHO", zap.String("请求", string(p)))
	return len(p), nil
}
