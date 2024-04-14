package router

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	authhandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/authHandler"
	employerhandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/employerHandler"
	rolehandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/roleHandler"
	taghandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/tagHandler"

	logger "github.com/sirupsen/logrus"
	// "github.com/AbdulfatahMohammedSheikh/backend/middlewares/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(
			200, gin.H{"message": "up and running"},
		)
	})

	authhandler.Regiester(r, a, log)
	taghandler.Regiester(r, a, log)
	rolehandler.Regiester(r, a, log)
	employerhandler.Regiester(r, a, log)
	// TODO: add the fallowing handlers
	// homehandler -> for browsing the site
	// user hanlder
	// employer handler

}
