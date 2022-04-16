package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/rd67/go-gin-mySQL/configs"
	"github.com/rd67/go-gin-mySQL/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
)

func dbURLGenerate() string {
	dbConfig := configs.GetConfig().Db

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// res2B, _ := json.Marshal(dbConfig)
	// fmt.Println(string(res2B))

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}

func ConnectDb() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var dsn = dbURLGenerate()
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// NamingStrategy: schema.NamingStrategy{
		// 	NoLowerCase: true,
		// },
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("DB Connected")
	db.AutoMigrate(&models.User{}, &models.Team{}, &models.TeamUser{}, &models.TeamUserStep{})
	DBConn = db
}
