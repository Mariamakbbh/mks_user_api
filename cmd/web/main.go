package main

import (
	"log"

	"github.com/Mariamakbbh/mks_LoginAPI/internal/config"
	"github.com/Mariamakbbh/mks_LoginAPI/internal/db"
	"github.com/Mariamakbbh/mks_LoginAPI/pkg/userLogin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(&cfg)

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	repo := userLogin.NewRepo(db)

	svc := userLogin.NewService(repo)

	bluePrint := userLogin.NewBlueprint(svc)

	bluePrint.Routes(app)

	app.Listen(cfg.Port)

	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// // Block until we receive our signal.
	// <-c

	// // Doesn't block if no connections, but will otherwise wait
	// // until the timeout deadline.
	// // Optionally, you could run srv.Shutdown in a goroutine and block on
	// // <-ctx.Done() if your application should wait for other services
	// // to finalize based on context cancellation.
	// log.Printf("Shutting down")
	// os.Exit(0)
}
