package app

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
)

func StartApp() {
	// Load env variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init gin
	router = gin.Default()

	// Init Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	router.Use(sentrygin.New(sentrygin.Options{}))

	// Bind custom validation
	bindCustomRules()

	// Setup application routes
	setUpRoutes()

	// Run server
	router.Run(":" + os.Getenv("PORT"))
}
