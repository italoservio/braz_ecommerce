package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/italoservio/braz_ecommerce/internal/create_user"
	"github.com/italoservio/braz_ecommerce/packages/database"
)

type Healthcheck struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

func main() {
	app := fiber.New()

	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	db, err := database.NewDatabase(dbUri, dbName)
	if err != nil {
		log.Fatal(err)
	}

	crudRepository := database.NewCrudRepository(db)
	createUserRepo := create_user.NewUserRepository(db, crudRepository)
	createUserService := create_user.NewUserService(createUserRepo)
	createUserController := create_user.NewUserController(createUserService)

	app.Post("/user", createUserController.CreateUser)
	app.Get("/health", healthcheck(db))

	go func() {
		log.Fatal(app.Listen(":3000"))
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
	db.Client().Disconnect(ctx)
}

func healthcheck(db *database.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		healthCheck := Healthcheck{
			Server:   "healthy",
			Database: "healthy",
		}

		err := db.Client().Ping(ctx, nil)

		if err != nil {
			healthCheck.Database = "offline"
		}

		return c.JSON(healthCheck)
	}
}
