//go:build wireinject

package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"log"
	"net/http"
	"os"
	"time"
)

func setupLogging() *Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix

	return &Logger{
		logInfo:    log.New(os.Stdout, "INFO: ", flags),
		logError:   log.New(os.Stdout, "ERROR: ", flags),
		logWarning: log.New(os.Stdout, "WARN: ", flags),
	}
}
func initializeConfig(l *Logger) (c Config) {
	var configLocation string
	flag.StringVar(&configLocation, "c", "config.toml", "Location of config file")
	flag.Parse()

	rawConfig, err := os.ReadFile(configLocation)
	if err != nil {
		l.logError.Fatal(err)
	}

	if _, err = toml.Decode(string(rawConfig), &c); err != nil {
		l.logError.Fatal(err)
	}

	return
}

func setupRatelimiting() *limiter.Limiter {
	rateLimiter := tollbooth.NewLimiter(2, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	rateLimiter.SetIPLookups([]string{"X-Forwarded-For", "RemoteAddr", "X-Real-IP"})

	return rateLimiter
}

func addRouter(uninitializedApp *uninitializedApplication) (app *Application) {
	app = (*Application)(uninitializedApp)
	app.logInfo.Println("Setting up router")
	gin.SetMode(gin.ReleaseMode)
	app.Router = gin.Default()

	api := app.Router.Group("/api")
	api.GET("/collectibles-account-value", app.collectiblesAccountValueAPI)
	api.GET("/can-view-inventory", app.canViewInventoryAPI)
	api.GET("/profile-info", app.profileInfoAPI)
	api.GET("/exchange-rate", app.exchangeRateAPI)
	api.Use(app.ratelimitMiddleware())

	app.Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/ayes-web/roblox-account-value-api")
	})

	return
}

type uninitializedApplication Application

func initializeApplication() *Application {
	panic(wire.Build(wire.NewSet(
		setupLogging,
		initializeConfig,
		setupRatelimiting,

		wire.Struct(
			new(uninitializedApplication),
			"Logger",
			"config",
			"RateLimiter",
		),

		addRouter,
	)))
}
