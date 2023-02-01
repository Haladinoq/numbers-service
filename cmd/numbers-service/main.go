package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/rs/zerolog"
	logg "github.com/rs/zerolog/log"
	_ "gitlab.palo-it.net/palo/numbers-service/docs"
	"gitlab.palo-it.net/palo/numbers-service/pkg/config"
	"gitlab.palo-it.net/palo/numbers-service/pkg/middleware"
	"gitlab.palo-it.net/palo/numbers-service/pkg/numbers"
	"gitlab.palo-it.net/palo/numbers-service/pkg/swagger"
	"gitlab.palo-it.net/palo/numbers-service/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	cFile = flag.String("config", "", "Configuration file YML.")
)

// @title           Numbers Service
// @version         1.0
// @description     This is a Numbers Service.

// @contact.name   Numbers Service Support
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logg.Logger = logg.Output(zerolog.ConsoleWriter{Out: os.Stderr}) //Pretty logging
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	container := container.New()
	flag.Parse()

	var config *config.Config
	container.Singleton(configFactory)
	container.Resolve(&config)

	var context utils.Context

	var router *gin.Engine
	container.Singleton(routerFactory)
	container.Resolve(&context)

	router = context.Router

	container.Singleton(GormFactory)

	container.Singleton(httpClientFactory)

	router.SetTrustedProxies(nil)
	middleware.Middleware(router)

	numbers.RunGeneralConfig(&container)
	swagger.InitSwagger(router, config)

	if err := router.Run(":" + strconv.Itoa(config.ServerPort)); err != nil {
		log.Fatal(err)
	}

}

func configFactory() *config.Config {
	// Load config file.
	conf, err := config.LoadConfig(*cFile)
	if err != nil {
		log.Fatal(err)
		panic("Error on ")
	}
	return conf
}

func httpClientFactory() *http.Client {
	return &http.Client{}
}
func routerFactory() utils.Context {
	return utils.Context{
		Router: gin.New(),
	}
}
func GormFactory(c *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
