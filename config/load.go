package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	Secret        string
	RefreshSecret string
	MT_Server_Key string
}

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_NAME string
	DB_HOST string
	DB_PORT int
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res  == nil {
		logrus.Fatal("Config : Cannot start program, Failed to load configuration")
		return nil
	}
	return res
}

func LoadDBConfig() DatabaseConfig {
	var res = new(DatabaseConfig)
	godotenv.Load(".env")
	if val, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(val)
        if err!= nil {
            logrus.Fatal(err)
        }
		res.DB_PORT = port
	}
	if val, found := os.LookupEnv("DB_HOST"); found {
        res.DB_HOST = val
    }
	if val, found := os.LookupEnv("DB_USER"); found {
        res.DB_USER = val
    }
	if val, found := os.LookupEnv("DB_PASS"); found {
        res.DB_PASS = val
    }
	if val, found := os.LookupEnv("DB_NAME"); found {
        res.DB_NAME = val
    }
	return *res
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
	if val,found := os.LookupEnv("MT_SERVER_KEY"); found {
		res.MT_Server_Key = val
	}

	return res
}