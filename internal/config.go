package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	MONGOURI string
}

func (env *EnvConfig) LoadConfig() error {

	err_env := godotenv.Load(".env")

	if err_env != nil {
		log.Fatalf(err_env.Error())
	}

	env.MONGOURI = os.Getenv("MONGOURI")

	return nil

}
