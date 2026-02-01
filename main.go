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
	"github.com/m-bromo/atom-ly/internal/repository"
	"github.com/m-bromo/atom-ly/internal/service"
	"github.com/m-bromo/atom-ly/internal/web/handler"
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

	c := gin.New()

	querier := sqlc.New(db)
	linkRepository := repository.NewPostgresLinkRepository(querier)
	hasher := hasher.NewHashID()
	linkService := service.NewLinkService(linkRepository, hasher)
	linkHandler := handler.NewLinkHandler(linkService)

	routes.SetupRoutes(c, linkHandler)

	log.Fatal(c.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port)))

	slog.Info("Starting application")
}
