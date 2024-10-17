package main

import (
	"SerialArduinoCommunication/configuration"
	"SerialArduinoCommunication/handlers"
	"SerialArduinoCommunication/initializers"
	"SerialArduinoCommunication/storage"
	"errors"
	"fmt"
	"github.com/easonlin404/limit"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	config := configuration.GetConfig()

	initializers.InitializeRedis()
	//initializers.InitializeEmailSender()

	err := storage.CreteConnection()
	if err != nil {
		log.Fatal(fmt.Sprintf("Server gone away >>> %v", err.Error()))

	}

	//migrate.MakeMigrations()

	gin.SetMode(config.Server.Mode)
	e := gin.New()

	// Prepare CORS
	e.Use(cors.Middleware(cors.Config{
		Origins:         strings.Join(config.Server.Origin.Cors, ", "),
		Methods:         "GET, POST, PUT, OPTIONS,DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Http-Auth, X-Last-Page, X-Device",
		ExposedHeaders:  "",
		MaxAge:          time.Duration(config.Server.Timeout.Read) * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	e.Use(limit.Limit(500))

	e.NoRoute(handlers.NoRouteHandler)

	handlers.HealthCheckHandler(e)
	handlers.LedHandler(e)
	handlers.RemoteHandler(e)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port),
		Handler:      e,
		ReadTimeout:  time.Duration(config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(config.Server.Timeout.Write) * time.Second,
	}

	// Run server inside goroutine that will not block signal
	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("Server gone away >>> %v", err.Error()))
	}

}
