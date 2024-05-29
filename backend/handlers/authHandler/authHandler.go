package authhandler

import (
	"fmt"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	roleroutermiddleware "github.com/AbdulfatahMohammedSheikh/backend/middlewares/role_router_middleware"
	authmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/authModal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
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
			c.String(401, "check the information you entered -----")
			return
		}

		tokenString, role, id, err := authmodal.Login(a, req.Email, req.Password)

		if err != nil {
			c.JSON(401, gin.H{
				"error": "cannot create token",
			})
			return
		}
		c.Set("token", *tokenString)
		c.Set("role", *role)
		c.Set("id", *id)
		c.Next()

		// //creating cookes
		// c.SetSameSite(http.SameSiteLaxMode)

		// // TODO: store the token in server and compare it later

		// c.SetCookie("authtoken", *tokenString, 3600*20*30, "/", "http://127.0.0.1:8080", true, true)
		// c.SetCookie("token", "dd", 606024, "/", "http://localhost", true, true)

		// c.JSON(200, gin.H{
		// 	"key": tokenString,
		// })
	}, roleroutermiddleware.RoleRouter)

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

		role, err := rolemodal.HasRoleWithName(a, req.Role)

		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

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

		c.JSON(200, gin.H{})
	})

	// signout
	r.POST("/auth/signout", func(c *gin.Context) {
		// remove the tokens from cookies
		// then redirect use to log in page
	})
}
