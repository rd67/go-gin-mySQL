package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfigStruct struct {
	Name    string
	Port    string
	BaseURL string
	AppURL  string
}

type DbConfigStruct struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
	Timezone string
}

type ConfigStruct struct {
	App AppConfigStruct
	Db  DbConfigStruct
}

func init() {
	//This func runs before the main func
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
}

func GetConfig() *ConfigStruct {

	var AppConfig = AppConfigStruct{
		Name:    os.Getenv("APP_NAME"),
		Port:    os.Getenv("PORT"),
		BaseURL: os.Getenv("BASE_URL"),
		AppURL:  os.Getenv("APP_URL"),
	}

	var DbConfig = DbConfigStruct{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}

	var Config = ConfigStruct{
		App: AppConfig,
		Db:  DbConfig,
	}

	return &Config
}
