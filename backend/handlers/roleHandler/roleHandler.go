package rolehandler

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	rolemodal "github.com/AbdulfatahMohammedSheikh/backend/modals/roleModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {
	// create
	r.POST("/role", func(c *gin.Context) {
        // TODO: did not test this 

		var req struct {
			Name string `form:"name" binding:"required"`
		}
		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "check the information you entered",
			})
			return
		}
		role := rolemodal.New(req.Name)
		err = role.Create(a)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	// index -> will return all the user with given role
	r.GET("/role", func(c *gin.Context) {
		id := c.Query("id")
		if "" == id {
			c.JSON(401, gin.H{
				"error": "id was not given",
			})
			return
		}

		_, err := rolemodal.HasRoleWithId(a, id)
		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		r, users, err := usermodal.Roles(a, id)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		if 0 == len(users) {
			c.JSON(401, gin.H{
				"error": "no users with given role",
			})
			return
		}

		if "" == r.Name {
			c.JSON(401, gin.H{
				"error": "role has no name",
			})
			return
		}

		c.JSON(200, gin.H{
			"role":  r,
			"users": users,
		})
	})

	// get all role
	r.GET("/roles", func(c *gin.Context) {
		roles := rolemodal.All(a)

		c.JSON(200, gin.H{
			"roles": roles,
		})
	})

}
