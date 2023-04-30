package main

import (
	"github.com/didip/tollbooth/v7/limiter"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Application struct {
	*Logger
	RateLimiter *limiter.Limiter
	Router      *gin.Engine

	config Config
}

type Config struct {
	RobuxToEuroRate uint64 `toml:"robux_to_euro_rate"`
	Port            string `toml:"port"`
}

func (app *Application) run() {
	app.logInfo.Printf("Starting server at http://localhost:%s\n", app.config.Port)
	app.logError.Fatal(http.ListenAndServe(":"+app.config.Port, app.Router))
}

type Logger struct {
	logError   *log.Logger
	logWarning *log.Logger
	logInfo    *log.Logger
}

func main() {
	app := initializeApplication()
	app.run()
}
