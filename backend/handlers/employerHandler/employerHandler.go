package employerhandler

import (
	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	employermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/employerModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	r.GET("/ngo/add-offer", func(c *gin.Context) {
		// TODO: implement this

		c.JSON(200, gin.H{})
	})

	r.GET("/ngo/update-offer/:id", func(c *gin.Context) {

		// TODO: implement this

		c.JSON(200, gin.H{})

	})

	// create
	r.POST("/employer", func(c *gin.Context) {
		var req struct {
			Name     string `form:"name" binding:"required"`
			UserId   string `form:"user_id" binding:"required"`
			Meto     string `form:"meto" binding:"required"`
			About    string `form:"about" binding:"required"`
			Location string `form:"location" binding:"required"`
			Phone    string `form:"phone" binding:"required"`
		}

		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "check the information you entered",
			})
			return
		}

		employer := employermodal.New(req.UserId, req.Name, req.Meto, req.About, req.Location, req.Phone)
		err = employer.Create(a)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	r.GET("/employer", func(c *gin.Context) {
		id := c.Query("id")
		if "" == id {
			c.JSON(404, gin.H{
				"error": "id was not provided",
			})
			return
		}

		emploter, err := employermodal.Index(a, id)
		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		if "" == emploter.Id {
			c.JSON(404, gin.H{
				"error": "there is no id in the return result check the operation of creating employer",
			})
			return

		}

		c.JSON(200, gin.H{})
	})

	r.GET("/employers", func(c *gin.Context) {
		employers, err := employermodal.All(a)
		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return

		}

		c.JSON(200, gin.H{
			"employers": employers,
		})
	})

	r.PATCH("/employer", func(c *gin.Context) {

		var req struct {
			Id       string `form:"id" binding:"required"`
			Name     string `form:"name" binding:"required"`
			UserId   string `form:"user_id" binding:"required"`
			Meto     string `form:"meto" binding:"required"`
			About    string `form:"about" binding:"required"`
			Location string `form:"location" binding:"required"`
			Phone    string `form:"phone" binding:"required"`
		}

		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "check the information you entered",
			})
			return
		}

		employer := employermodal.Employer{
			Id:       req.Id,
			Name:     req.Name,
			UserId:   req.UserId,
			Meto:     req.Meto,
			Location: req.Location,
			About:    req.About,
			Phone:    req.Phone,
		}

		err = employermodal.Update(a, employer)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	r.POST("/delete/employer", func(c *gin.Context) {

		var req struct {
			Id string `form:"id" binding:"required"`
		}

		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "id was not provided",
			})
			return
		}
		err = employermodal.Delete(a, req.Id)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	r.POST("/employer/tag", func(c *gin.Context) {

		var req struct {
			Tag string `form:"tag" binding:"required"`
			Id  string `form:"id" binding:"required"`
		}

		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "check the information you entered",
			})
			return
		}
		err = employermodal.AddTag(a, req.Id, req.Tag)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	r.POST("/employer/delete/tag", func(c *gin.Context) {

		var req struct {
			Tag string `form:"tag" binding:"required"`
			Id  string `form:"id" binding:"required"`
		}

		err := c.ShouldBind(&req)
		if nil != err {
			c.JSON(404, gin.H{
				"error": "check the information you entered",
			})
			return
		}
		err = employermodal.RemoveTag(a, req.Id, req.Tag)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})
}
