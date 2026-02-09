package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/config"
	"github.com/m-bromo/atom-ly/internal/database/postgres"
	"github.com/m-bromo/atom-ly/internal/database/postgres/sqlc"
	"github.com/m-bromo/atom-ly/internal/hasher"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
	"github.com/m-bromo/atom-ly/internal/service"
	"github.com/m-bromo/atom-ly/internal/web/handler"
	"github.com/m-bromo/atom-ly/internal/web/middleware"
	"github.com/m-bromo/atom-ly/internal/web/routes"
	"github.com/m-bromo/atom-ly/logger"
)

func main() {
	config.SetupEnvironment()
	logger.SetupLog(config.Env.Environment)

	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	c := gin.Default()

	querier := sqlc.New(db)
	hasher := hasher.NewHashID()
	linkRepository := repository.NewPostgresLinkRepository(querier)
	linkService := service.NewLinkService(linkRepository, hasher)
	linkHandler := handler.NewLinkHandler(linkService)
	errorMidleware := middleware.NewErrorMiddleware()

	routes.SetupRoutes(c, linkHandler, errorMidleware)

	log.Fatal(c.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port)))

	slog.Info("Starting application")
}
