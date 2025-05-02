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

func GetEmail() string {
	return os.Getenv("EMAIL")
}

func GetSenha() string {
	return os.Getenv("SENHA")
}
