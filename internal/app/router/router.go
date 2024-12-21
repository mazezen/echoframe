package router

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/echoframe/internal/app/middlewares"
	"time"

	"github.com/labstack/echo/v4/middleware"
)

func BootApp(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.HEAD, echo.PUT, echo.POST, echo.GET, echo.OPTIONS, echo.PATCH, echo.DELETE},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 24,
		}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ip=${remote_ip} time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
		Output: middlewares.EchoLog,
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(middlewares.FrameLog())
	e.Use(middleware.BodyDumpWithConfig(middlewares.ConsoleBodyDumpConfig))

	//e.Use(middles.ErrorHandlingMiddleware)
	//e.Use(middles.ResponseMiddleware())

}
