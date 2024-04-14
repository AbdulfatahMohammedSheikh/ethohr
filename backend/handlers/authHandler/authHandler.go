package authhandler

import (
	"net/http"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {
	// TODO: store tokens in server along with user id

	// signin
	r.POST("/auth/login", func(c *gin.Context) {

		var req struct {
			Email    string `form:"email" binding:"required"`
			Password string `form:"password" binding:"required"`
		}

		err := c.ShouldBind(&req)

		if nil != err {
			c.String(401, "check the information you entered")
			return
		}

		token, err := authmodal.Login(a, req.Email, req.Password)

		if nil != err {
			c.String(401, err.Error())
			return
		}

		// Sign and get the complete encoded token as a string using the secret
		m := make(map[string]string)
		m["k"] = "k"

		//TODO: the SignedString method cannot take string directly
		tokenString, err := token.SignedString([]byte("sfsd"))

		if err != nil {

			c.JSON(404, gin.H{
				"error": "cannot create token",
			})
			return
		}

		//creating cookes
		c.SetSameSite(http.SameSiteLaxMode)
		// TODO: store the token in server and compare it later
		c.SetCookie("authtoken", tokenString, 3600*20*30, "", "", true, true)
		c.JSON(200, gin.H{})
	})

	// signup
	r.POST("/auth/signup", func(c *gin.Context) {
		var req struct {
			Name     string `form:"name" binding:"required"`
			Email    string `form:"email" binding:"required"`
			Phone    string `form:"phone" binding:"required"`
			Password string `form:"password" binding:"required"`
			Role     string `form:"role" binding:"required"`
		}

		err := c.ShouldBind(&req)

		if nil != err {

			c.JSON(401, gin.H{
				"error": "check the information you entered",
			})
			return
		}
		err = authmodal.SignUp(a, req.Name, req.Email, req.Password, req.Phone, req.Role)

		if nil != err {

			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	// signout
	r.POST("/auth/signout", func(c *gin.Context) {
		// remove the tokens from cookies
		// then redirect use to log in page
	})
}
