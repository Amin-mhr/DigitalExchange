package api

import (
	"DigitalExchange/src/api/http/middlewares"
	"DigitalExchange/src/api/routes"
	"DigitalExchange/src/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	configs = config.GetInstance()
	g       errgroup.Group
)

func Init() (err error) {
	g.Go(func() error {
		return initUserServer()
	})

	if err = g.Wait(); err != nil {
		log.Fatalln(err)
		return err
	}

	return err
}

func getNewRouter() *gin.Engine {

	// set gin to release mode.
	gin.SetMode(gin.ReleaseMode)

	// Initialize new app.
	router := gin.New()

	// Attach CORS middlewares.
	router.Use(middlewares.Cors())

	// Attach logger middlewares.
	router.Use(gin.Logger())

	// Attach recovery middlewares.
	router.Use(gin.Recovery())

	// Attach Global Rate Limiter
	limiter := middlewares.NewRateLimiter(0.5, 3)
	router.Use(limiter.Middleware())

	return router
}

func initUserServer() error {
	router := getNewRouter()

	v1 := router.Group("api/v1")
	{
		routes.RegisterExchangeRoutes(v1)
	}

	// Run App.
	if err := router.Run(
		fmt.Sprintf(":%s", configs.Get("APP_PORT")),
	); err != nil {
		return err
	}

	return nil
}
