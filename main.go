package envUtil

import (
	"log"

	"github.com/joho/godotenv"
)

type EnvLoader interface{
	Getenv(str string)(string)
}

func GetEnv(env string, envLoader EnvLoader ) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading ENV")
	}
	variable := envLoader.Getenv(env)

	if variable == "" {
		log.Fatal("Failed to find env variable")
	}

	return variable
}
