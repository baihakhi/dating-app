package main

import (
	"fmt"
	"log"
	"os"

	driver "github.com/baihakhi/dating-app/internal/databases"
	"github.com/baihakhi/dating-app/internal/handler"
	"github.com/baihakhi/dating-app/internal/repositories"
	"github.com/baihakhi/dating-app/internal/router"
	"github.com/baihakhi/dating-app/internal/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Println(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	e := echo.New()
	e.Use(echomiddleware.Logger())
	e.Logger.SetLevel(2)

	SetupMiddleware(e)

	repositories := repositories.InitRepository(driver.Config())
	services := services.InitService(repositories)
	handlers := handler.InitiHandler(services)

	router.InitRouter(e, handlers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func SetupMiddleware(e *echo.Echo) {
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"Origin",
			"Authorization",
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Accept",
			"Content-Length",
			"Accept-Encoding",
		},
	}))
}
