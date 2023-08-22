package app

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Firgisotya/go-commerce/database/seeders"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB 		*gorm.DB
	Router 	*mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBDriver string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to" + appConfig.AppName)

	
	server.initiaLizeRoutes()
}

func (server *Server) initializeDB(dbConfig DBConfig){
	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := dbConfig.DBUser + ":" + dbConfig.DBPass + "@tcp(" + dbConfig.DBHost + ":" + dbConfig.DBPort + ")/" + dbConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		// postgresql
		dsn := dbConfig.DBUser + ":" + dbConfig.DBPass + "@tcp(" + dbConfig.DBHost + ":" + dbConfig.DBPort + ")/" + dbConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbConfig.DBDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", dbConfig.DBDriver)
	}

	

}

func (server *Server) dbMigrate(){
	for _, model := range RegisterModel() {
		err := server.DB.AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully")
}

func (server *Server) initCommands(config AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}


func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	// Get app config
	appConfig.AppName = getEnv("APP_NAME", "Go Commerce")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	// get db config
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBPort = getEnv("DB_PORT", "3306")
	dbConfig.DBUser = getEnv("DB_USER", "root")
	dbConfig.DBPass = getEnv("DB_PASS", "")
	dbConfig.DBName = getEnv("DB_NAME", "go_commerce")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "mysql")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.initCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}