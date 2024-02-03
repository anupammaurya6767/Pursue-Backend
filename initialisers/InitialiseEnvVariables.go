package initialisers

import (
	"fmt"
	"github.com/joho/godotenv"
)

// InitialiseEnvVariables: A method to initialise Environment variables
func InitialiseEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}
