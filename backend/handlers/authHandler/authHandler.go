package authhandler

import (
	"fmt"
	"net/http"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	// TODO: find a way to store the tokens in server or making sure the token is authantic

	// signin
	r.POST("/auth/login", func(c *gin.Context) {

		var data = map[string]interface{}{}

		var req struct {
			Email    string `form:"email" binding:"required"`
			Password string `form:"password" binding:"required"`
		}

		err := c.ShouldBind(&req)

		if nil != err {
			data["error"] = err.Error()
			c.HTML(422, "error", data)
			return
		}

		tokenString, role, id, err := authmodal.Login(a, req.Email, req.Password)

		if err != nil {
			data["error"] = err.Error()
			c.HTML(422, "error", data)
			return
		}

		// TODO: delete the uncorrent token data

		// //creating cookes
		c.SetSameSite(http.SameSiteLaxMode)

		// // TODO: store the token in server and compare it later

		c.SetCookie("authtoken", *tokenString, 3600*20*30, "/", "http://127.0.0.1:8080", true, true)
		c.SetCookie("id", *id, 3600*20*30, "/", "http://127.0.0.1:8080", true, true)

        // TODO: remove this black
		c.Set("token", *tokenString)
		c.Set("role", *role)
		c.Set("id", *id)

		// c.SetCookie("token", "dd", 606024, "/", "http://localhost", true, true)

		// c.JSON(200, gin.H{
		// 	"key": tokenString,
		// })
		// TODO: make this smart that it return the home page for each role

		c.Header("HX-Redirect", "/nog-home")

		// c.JSON(200, gin.H{})

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

		role, err := rolemodal.HasRoleWithId(a, req.Role)

		if nil != err {

			c.JSON(401, gin.H{
				"error": "You are trying to create accound for nonexiting role",
			})
			return
		}

		err = authmodal.SignUp(a, req.Name, req.Email, req.Password, req.Phone, role.Id)

		if nil != err {

			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Header("HX-Redirect", "/login")
		c.JSON(200, gin.H{})

	})

    // TODO: implement this
	// signout
	r.POST("/auth/signout", func(c *gin.Context) {
		// remove the tokens from cookies
		// then redirect use to log in page
	})
}
