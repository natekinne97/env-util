package envUtil

import (
	"log"
	"testing"
)

type FakeEnvLoader struct{}

func (_ FakeEnvLoader) Getenv(str string)(string){
	return str
}

func TestGetEnv(t *testing.T){
	envLoader := FakeEnvLoader{}

	envVar := GetEnv("test", envLoader)

	if envVar == "test"{
		log.Println("Successfully retrieved env var test")
	}else{
		log.Println("Failed to get env var")
	}
}
