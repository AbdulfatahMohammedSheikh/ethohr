package offerhandler

import (

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	offermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/offerModal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Regiester(r *gin.Engine, a *surreal.AppRepository, log *logger.Logger) {

	r.POST("/offer", func(c *gin.Context) {
		// TODO: use the user_id to get the employer name

		var req struct {
			EmployerId string `form:"employer_id" binding:"required"`
			// EmployerName string `form:"employer_name" binding:"required"`
			Title        string `form:"title" binding:"required"`
			Requirements string `form:"requirements" binding:"required"`
			Duty         string `form:"duty" binding:"required"`
			PostDate     string `form:"postDate" binding:"required"`
			Deadline     string `form:"deadline" binding:"required"`
		}
		err := c.ShouldBind(&req)

		if nil != err {
			c.JSON(401, gin.H{
				"error": "check the information you entered",
			})
			return
		}

		user, err := usermodal.Index(a, req.EmployerId)
		_ = user

		if nil != err {
			c.JSON(401, gin.H{
				"error": "error getting the name of the ngo",
			})
			return
		}

		offer := offermodal.New(
			req.EmployerId,
			// TODO: make the below line use the infratmion from where employer create his name
			req.EmployerId,
			req.Title,
			req.Requirements,
			req.Duty,
			req.PostDate,
			req.Deadline,
		)

		err = offer.Create(a)

		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{})
	})

	r.GET("/offer", func(c *gin.Context) {
		ngo := c.Query("ngo")
		if "" == ngo {
			c.JSON(401, gin.H{
				"error": "did not provide ngo id",
			})
			return
		}

		offers, err := offermodal.Index(a, ngo)

		if nil != err {
			c.JSON(401, gin.H{
				"error": "did not provide ngo id",
			})
			return
		}

		c.JSON(200, gin.H{
			"offers": offers,
		})
	})

	r.GET("/offer/show/:id", func(c *gin.Context) {
		id := c.Param("id")
		if "" == id {
			c.JSON(401, gin.H{
				"error": "id was not provided",
			})
			return
		}

		offer, err := offermodal.ShowOffer(a, id)
		if nil != err {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"offers": offer,
		})
	})

}
