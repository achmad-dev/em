package v1

import (
	"fmt"

	"github.com/achmad/em/backend/api/handler"
	v1 "github.com/achmad/em/backend/api/middleware/v1"
	"github.com/achmad/em/backend/config"
	"github.com/achmad/em/backend/internal/repository"
	"github.com/achmad/em/backend/internal/service"
	"github.com/achmad/em/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"golang.org/x/crypto/bcrypt"
)

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

func ServeRoute(envPath string) {
	log := utils.InitLog()
	log.Info("Route API v1 is running")
	cfg, err := config.NewConfig(envPath)
	if err != nil {
		log.Fatal("Failed to load configuration", err)
	}

	// Initialize the sqlx database
	postgreUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := utils.InitSqlDB(postgreUrl)
	if err != nil {
		log.Fatalf("failed to connect to database: %v url:%s", err, postgreUrl)
	}

	// initialize the services
	bcryptUtil := utils.NewBcryptUtil(bcrypt.DefaultCost)
	// initialize the user service
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, bcryptUtil, cfg.JwtSecret, log)

	// initialize the user handler
	userHandler := handler.NewAuthHandler(userService)

	// initialize event service
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo, userRepo, log)

	//initialize request log
	reqRepo := repository.NewRequestLogRepository(db)
	reqService := service.NewRequestLogService(reqRepo, log)

	// initialize event handler
	eventHandler := handler.NewEventHandler(eventService, userService, reqService)

	// initialize the fiber app
	app := fiber.New(
		fiber.Config{
			StrictRouting: true,
			Prefork:       true,
			AppName:       "EM",
		},
	)
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
		}),
		fiberLog.New(
			fiberLog.Config{
				Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
				TimeFormat: "02-Jan-2006",
			},
		),
	)

	// Define the routes
	api := app.Group("/api/v1", v1.HMACMiddleware(cfg.HmacSecret))
	api.Post("/signin", userHandler.SignIn)

	// Event routes
	event := api.Group("/events")
	event.Get("/companies", v1.AuthMiddleware(cfg.JwtSecret, userService), eventHandler.GetCompanies)
	event.Post("/insert", v1.AuthMiddleware(cfg.JwtSecret, userService), eventHandler.InsertEvent)
	event.Get("", v1.AuthMiddleware(cfg.JwtSecret, userService), eventHandler.GetEvents)
	event.Put("/update", v1.AuthMiddleware(cfg.JwtSecret, userService), eventHandler.UpdateEvent)

	// Metrics
	app.Get("/metrics", monitor.New(
		monitor.Config{
			Title: "Backend Metrics",
		},
	))

	// Start the server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.Port)))
}
