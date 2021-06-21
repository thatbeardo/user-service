// @title Fan fit login
// @version 0.1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @description ## Users
//@description ---
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @BasePath /
package main

import (
	"fmt"
	"os"
	"time"

	// API Routes
	clientHandlers "github.com/fanfit/login/api/handlers/clients"
	creatorHandlers "github.com/fanfit/login/api/handlers/creators"
	userHandlers "github.com/fanfit/login/api/handlers/users"

	// Tags
	// Users Tag
	userRepository "github.com/fanfit/login/models/users/repository"
	userServicePackage "github.com/fanfit/login/models/users/service"

	// Creators Tag
	creatorRepository "github.com/fanfit/login/models/creators/repository"
	creatorServicePackage "github.com/fanfit/login/models/creators/service"

	// Clients Tag
	clientRepository "github.com/fanfit/login/models/clients/repository"
	clientServicePackage "github.com/fanfit/login/models/clients/service"

	"github.com/fanfit/login/api/middleware"
	"github.com/fanfit/login/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	}

	db, err := server.CreatePostGresConnection(log, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Print("Something went wrong!" + err.Error())
	}

	// Instantiate service for each tag
	userStore := userRepository.NewUserStore(db)
	userService := userServicePackage.New(userStore)

	creatorStore := creatorRepository.NewUserStore(db)
	creatorService := creatorServicePackage.New(creatorStore)

	clientStore := clientRepository.NewUserStore(db)
	clientService := clientServicePackage.New(clientStore)

	// Initialize the middleware and routes
	engine := gin.Default()
	router := server.GenerateRouter(engine)

	// Set routes for each tag
	router.Use(middleware.VerifyToken)
	userHandlers.Routes(router, userService)
	clientHandlers.Routes(router, clientService)
	creatorHandlers.Routes(router, creatorService)

	server.Orchestrate(engine, db)
}
