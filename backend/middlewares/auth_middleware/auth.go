package authmiddleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/AbdulfatahMohammedSheikh/backend/db/surreal"
	usermodal "github.com/AbdulfatahMohammedSheikh/backend/modals/userModal"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func Auth(c *gin.Context) {
	tok, err := c.Cookie("authkey")

	if nil != err {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// TODO: use .env to store the app secrete key
		// creating secrete key for jwt
		return []byte("key"), nil
	})

	if nil != err {

		log.Fatal(err)
		// TODO: make the server return a html page instead of returing string
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check if the cookie is experied

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		config := surreal.NewApp()

		repo, err := surreal.NewAppRepository(config.DB)

		if nil != err {
			log.Fatalf("failed to creat app : %v", err)
		}

		log.Info("connecting to database ")
		defer func() {
			repo.Close()
		}()
		m := map[string]interface{}{}

		u, err := surreal.Find[usermodal.User](
			repo,
			"select * from users where id = $id",
			m,
			[]usermodal.User{},
		)

		if u[0].Id == "" {

			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return

		}
		c.Set("u", u[0])
		c.Next()
	} else {
		log.Fatal(err)
	}
}
