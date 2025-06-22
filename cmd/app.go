package cmd

import (
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
)

type Application struct {
	*Logger
	RateLimiter *limiter.Limiter
	Router      *gin.Engine

	config Config
}

type Config struct {
	RobuxPerEuro uint64 `toml:"robux_per_euro"`
	Port         string `toml:"port"`
}

func (app *Application) Run() {
	app.logInfo.Printf("Starting server at http://localhost:%s\n", app.config.Port)
	app.logError.Fatal(http.ListenAndServe(":"+app.config.Port, app.Router))
}

type Logger struct {
	logError   *log.Logger
	logWarning *log.Logger
	logInfo    *log.Logger
}
