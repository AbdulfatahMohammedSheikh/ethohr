package router

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	authhandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/authHandler"
	employerhandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/employerHandler"
	offerhandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/offerHandler"
	rolehandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/roleHandler"
	taghandler "github.com/AbdulfatahMohammedSheikh/backend/handlers/tagHandler"

	logger "github.com/sirupsen/logrus"
	// "github.com/AbdulfatahMohammedSheikh/backend/middlewares/auth"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400,
			gin.H{
				"error": "no such route",
			})

	})



	r.GET("/health", func(c *gin.Context) {
		c.JSON(
			200, gin.H{"message": "up and running"},
		)
	})

	authhandler.Regiester(r, a, log)
	taghandler.Regiester(r, a, log)
	rolehandler.Regiester(r, a, log)
	employerhandler.Regiester(r, a, log)
	offerhandler.Regiester(r, a, log)


	// TODO: use this to create cookies
	// r.GET("/setcookie", func(c *gin.Context) {
	// 	cookie := http.Cookie{
	// 		Name:  "myCookie",
	// 		Value: "someValue",
	// 		Path:  "/",
	// 		// Set Secure to false for local testing
	// 		Secure: false,
	// 		// Set Domain to localhost for local testing
	// 		Domain:   "127.0.0.1",
	// 		HttpOnly: false,
	// 		MaxAge:   3600, // Expires in 1 hour
	// 	}
	// 	http.SetCookie(c.Writer, &cookie)
	// 	c.JSON(http.StatusOK, gin.H{"message": "Cookie set successfully!"})
	// })

}
