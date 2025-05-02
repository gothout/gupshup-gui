package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Carrega variáveis de ambiente
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis de ambiente do sistema.")
	}
}

func GetDatabaseURL() string {
	return os.Getenv("EMAIL")
}

func GetDatabaseUser() string {
	return os.Getenv("SENHA")
}
