package config


import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	InitDBEnvironmentVariables()
	InitDBConnection()
	InitDBCollections()
	echoInstance := echo.New()
	echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	echoInstance.Use(middleware.Recover())
	return echoInstance
}

