package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res  == nil {
		logrus.Fatal("Config : Cannot start program, Failed to load configuration")
		return nil
	}
	return res
}

type ProgramConfig struct {
	Secret        string
	RefreshSecret string
}

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT int
}

func LoadDBConfig() DatabaseConfig {
	godotenv.Load(".env")

	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic(err)
	}

	return DatabaseConfig{
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: DB_PORT,
		DB_NAME: os.Getenv("DB_NAME"),
	}
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	err := godotenv.Load(".env")

	if err!= nil {
        logrus.Error("Config : Cannot load config file: ", err.Error())
    }

	if val,found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val,found := os.LookupEnv("REFSECRET"); found {
		res.RefreshSecret = val
	}

	return res
}