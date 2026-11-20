package main

import (
	"log"

	"github.com/Edgar0bsj/NEADTech/internal/platform/config"
	"github.com/Edgar0bsj/NEADTech/internal/platform/database"

	httpapi "github.com/Edgar0bsj/NEADTech/internal/http"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewMySQL(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	router := httpapi.NewRouter()

	log.Printf("Server running on port %s", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
