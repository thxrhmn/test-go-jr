package main

import (
	"fmt"
	"net/http"
	"os"
	"todo-list-thxrhmn/database"
	"todo-list-thxrhmn/pkg/postgresql"
	"todo-list-thxrhmn/routes"

	_ "todo-list-thxrhmn/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5000
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	postgresql.DatabaseInit()
	database.RunMigration()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.RouteInit(e.Group("/api/v1"))

	var PORT = os.Getenv("PORT")

	fmt.Println("server running localhost:" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT))
}
