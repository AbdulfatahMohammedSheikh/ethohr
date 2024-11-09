package offerhandler

import (
	"fmt"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	offermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/offerModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	r.POST("/offer", func(c *gin.Context) {

		var req struct {
			EmployerId string `form:"employer_id" binding:"required"`
			// EmployerName string `form:"employer_name" binding:"required"`
			Title        string   `form:"title" binding:"required"`
			Requirements []string `form:"requirements" binding:"required"`
			Duty         []string `form:"duty" binding:"required"`
			PostDate     string   `form:"postDate" binding:"required"`
			Deadline     string   `form:"deadline" binding:"required"`
		}
		err := c.ShouldBind(&req)

		if nil != err {
			fmt.Println(err)
			return
		}

		user, err := usermodal.Index(a, req.EmployerId)
		_ = user

		if nil != err {
			fmt.Println(err)
			return
		}

		offer := offermodal.New(
			req.EmployerId,
			// TODO: make the below line use the infratmion from where employer create his name
			req.EmployerId,
			req.Title,
			req.PostDate,
			req.Deadline,
			req.Requirements,
			req.Duty,
		)

		err = offer.Create(a)

		if nil != err {

			fmt.Println(err)

			return
		}

		// c.Header("HX-Redirect", "/nog-home")
	})

	r.GET("/offers", func(c *gin.Context) {
		ngo := c.Query("ngo")

		ngoId, err := c.Cookie("id")

		data := map[string]interface{}{}
		data["title"] = "ngo offers"

		if nil != err {
			c.JSON(401, gin.H{})
			return
		}

		if "" == ngo {
			// TODO: deal with this case by rendering menaingful error massage
			c.JSON(401, gin.H{})
			return
		}

		if ngo != ngoId {
			c.JSON(401, gin.H{})
			return
		}

		offers, err := offermodal.Index(a, ngo)

		if nil != err {

			c.JSON(401, gin.H{})
			return
		}

		c.JSON(200, gin.H{
			"offers": offers,
		})

	})

	r.GET("/offer/show/:id", func(c *gin.Context) {

		id := c.Param("id")

		if "" == id {

			c.JSON(401, gin.H{})
			return
		}

		offer, err := offermodal.ShowOffer(a, id)
		if nil != err {

			c.JSON(401, gin.H{})
			return
		}

		c.JSON(200, gin.H{
			"offer": offer,
		})

	})

	r.DELETE("/offer/:id", func(c *gin.Context) {

		id := c.Param("id")
		from, err := c.Cookie("id")

		if id == "" {

			c.JSON(422, gin.H{})
			return
		}

		if nil != err {

			c.JSON(422, gin.H{
				"error": err.Error(),
			})

			return
		}

		err = offermodal.Delete(a, id, from)

		if nil != err {

			c.JSON(422, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(200, gin.H{})
	})

	r.PATCH("/offer/:id", func(c *gin.Context) {

		var req struct {
			EmployerId string `form:"employer_id" binding:"required"`
			// EmployerName string `form:"employer_name" binding:"required"`
			Title        string   `form:"title" binding:"required"`
			Requirements []string `form:"requirements" binding:"required"`
			Duty         []string `form:"duty" binding:"required"`
			PostDate     string   `form:"postDate" binding:"required"`
			Deadline     string   `form:"deadline" binding:"required"`
		}
		err := c.ShouldBind(&req)

		if nil != err {

			c.JSON(404, gin.H{
				"error": err.Error(),
			})

			return
		}

		user, err := usermodal.Index(a, req.EmployerId)
		_ = user

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		offer := offermodal.New(
			req.EmployerId,
			// TODO: make the below line use the infratmion from where employer create his name
			req.EmployerId,
			req.Title,
			req.PostDate,
			req.Deadline,
			req.Requirements,
			req.Duty,
		)


		err = offermodal.Update(a, *offer)

		if nil != err {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

}
