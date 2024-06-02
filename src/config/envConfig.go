package initializers

import (
	"fmt"
	"log"

	"github.com/lpernett/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	fmt.Println("Env Loaded Successfully...")
}
