package main

import (
	"github.com/AbdulfatahMohammedSheikh/backend/core/router"
	surreal "github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func NewLogger() *logger.Logger {
	return logger.New()
}

func main() {

	config := surreal.NewApp()

    // TODO: rewrite test
	// TODO: add auth using jwt


	var log = NewLogger()
	log.Trace(logger.TraceLevel)
	repo, err := surreal.NewAppRepository(config.DB)

	if nil != err {
		log.Fatalf("failed to creat app : %v", err)
	}

	log.Info("connecting to database ")
	defer func() {
		repo.Close()
	}()

	r := gin.Default()
	// corsConfig := cors.DefaultConfig()
	//
	// corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	// corsConfig.AllowMethods = []string{"PUT", "PATCH", "POST", "GET"}
	//
	// // Configure CORS using default settings (modify as needed)
	// r.Use(cors.New(corsConfig))
	router.SetRouter(r, repo, log)
	r.Run()
}
