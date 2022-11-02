package db

//https://levelup.gitconnected.com/api-in-go-with-fiber-and-docker-5de04651463a
import (
	"fmt"
	"log"

	"github.com/Mariamakbbh/mks_LoginAPI/internal/config"
	model "github.com/Mariamakbbh/mks_LoginAPI/pkg/userLogin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:root@%s:%s/%s", c.DBUser, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.UserDetails{})

	return db
}
