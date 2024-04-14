package taghandler

import (
	"fmt"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	tagmodal "github.com/AbdulfatahMohammedSheikh/backend/modals/tagModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	// var req struct {
	// 	Email    string `form:"email" binding:"required"`
	// 	Password string `form:"password" binding:"required"`
	// }
	//
	// err := c.ShouldBind(&req)

	// create
	r.POST("/tag", func(c *gin.Context) {
		var req struct {
			Name string `form:"name" binding:"required"`
		}
		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(401, gin.H{
				"error": "check the information you entered",
			})
			return
		}
		err = tagmodal.Create(a, req.Name)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"info": "tag created",
		})
	})

	// all
	r.GET("/tags", func(c *gin.Context) {
		tags, err := tagmodal.All(a)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"tags": tags,
		})
	})

	// index
	r.GET("/tag", func(c *gin.Context) {
		id := c.Query("id")

		if "" == id {
			c.JSON(401, gin.H{
				"error": "id is not provided",
			})
			return
		}

		tag, err := tagmodal.Index(a, id)

		if nil != err {
			if nil != err {
				c.JSON(401, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		c.JSON(200, gin.H{
			"tag": tag,
		})
	})

	// update

	r.PATCH("/tag", func(c *gin.Context) {
		var req struct {
			Name string `form:"name" binding:"required"`
			Id   string `form:"id" binding:"required"`
		}
		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(401, gin.H{
				"error": "check the information you entered",
			})
			return
		}

		tag := tagmodal.Tag{
			Id:   req.Id,
			Name: req.Name,
		}
		err = tagmodal.Update(a, tag)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"info": "tag updated",
		})
	})

	// delete

	r.POST("delete/tag", func(c *gin.Context) {
		var req struct {
			Id string `form:"id" binding:"required"`
		}
		err := c.ShouldBind(&req)
        fmt.Println("=================================----------------------------================================")
        fmt.Println(req.Id)
        fmt.Println("=================================----------------------------================================")
		if nil != err {
			c.JSON(401, gin.H{
				"error": "check the information you entered",
			})
			return
		}

		err = tagmodal.Delete(a, req.Id)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"info": "tag updated",
		})
	})
}
