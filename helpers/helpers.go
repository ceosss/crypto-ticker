package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetApiKey() string {
	godotenv.Load(".env")
	apiKey := os.Getenv("NOMICS_API_KEY")
	fmt.Println("NOMICS_API_KEY" + apiKey)
	return apiKey
}
