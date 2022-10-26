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
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	db := db.Init(&c)

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	repo := userLogin.NewRepo(db)

	svc := userLogin.NewService(repo)

	bluePrint := userLogin.NewBlueprint(svc)

	bluePrint.Routes(app)

	app.Listen(c.Port)
}
