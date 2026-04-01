package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/config"
	"github.com/m-bromo/atom-ly/internal/database/postgres"
	"github.com/m-bromo/atom-ly/internal/database/postgres/sqlc"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
	"github.com/m-bromo/atom-ly/internal/service"
	"github.com/m-bromo/atom-ly/internal/web/handler"
	"github.com/m-bromo/atom-ly/internal/web/middleware"
	resterrors "github.com/m-bromo/atom-ly/internal/web/rest_errors"
	"github.com/m-bromo/atom-ly/internal/web/routes"
	"github.com/m-bromo/atom-ly/pkg/hasher"
	"github.com/m-bromo/atom-ly/pkg/logger"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewPostgresConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	c := gin.Default()

	c.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowHeaders: []string{"Origin", "Content-type"},
	}))

	querier := sqlc.New(db)
	hasher := hasher.NewHashID(config)
	logger := logger.NewLogger(config)
	linkRepository := repository.NewPostgresLinkRepository(querier)
	linkService := service.NewLinkService(linkRepository, hasher)
	linkHandler := handler.NewLinkHandler(linkService, config)
	errHandler := resterrors.NewErrorHandler(logger)
	errorMidleware := middleware.NewErrorMiddleware(errHandler)

	routes.SetupRoutes(c, linkHandler, errorMidleware)

	log.Fatal(c.Run(fmt.Sprintf("%s:%s", config.Env.Api.Host, config.Env.Api.Port)))

	slog.Info("Starting application")
}
