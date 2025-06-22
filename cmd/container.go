//go:build wireinject

package cmd

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
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

func setupRatelimiting(c Config) *limiter.Limiter {
	rateLimiter := tollbooth.NewLimiter(2, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

	if c.BehindReverseProxy {
		rateLimiter.SetIPLookup(limiter.IPLookup{
			Name:           "X-Forwarded-For",
			IndexFromRight: 0,
		})
	} else {
		rateLimiter.SetIPLookup(limiter.IPLookup{
			Name:           "RemoteAddr",
			IndexFromRight: 0,
		})
	}

	return rateLimiter
}

func addRouter(uninitializedApp *uninitializedApplication) (app *Application) {
	app = (*Application)(uninitializedApp)
	app.logInfo.Println("Setting up router")
	app.Router = gin.Default()

	api := app.Router.Group("/api")
	api.GET("/collectibles-account-value", app.handleCollectiblesAccountValue)
	api.GET("/can-view-inventory", app.handleCanViewInventory)
	api.GET("/profile-info", app.handleProfileInfo)
	api.GET("/exchange-rate", app.handleExchangeRate)
	api.Use(app.ratelimitMiddleware())

	app.Router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://github.com/BatteredBunny/roblox-account-value-api")
	})

	return
}

type uninitializedApplication Application

func InitializeApplication() *Application {
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
