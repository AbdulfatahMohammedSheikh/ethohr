package roleroutermiddleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Routes       []string
	DefaultRoute string
}

func newRoleRouter(defaultRoute string, routes []string) *Router {
	return &Router{
		Routes:       routes,
		DefaultRoute: defaultRoute,
	}
}

func RoleRouter(c *gin.Context) {

	role, ok := c.Get("role")

	if !ok {
		c.JSON(401, gin.H{
			"error": "role is not settup",
		})
	}

	id, ok := c.Get("id")

	if !ok {
		c.JSON(401, gin.H{
			"error": "role is not settup",
		})
	}

	// TODO: uset this token later
	// token, ok := c.Get("token")

	switch role {

	// TODO: add user id
	// TODO: add the  role id
	case "ngo":

		defaultRoute := "/ngo/profile"
		routes := []string{
			"shit",
		}
		roleroutes := newRoleRouter(defaultRoute, routes)
		c.JSON(
			201,
			gin.H{
				"routes": roleroutes.Routes,
				"default_route": roleroutes.DefaultRoute,
                "id":id,
			},
		)

	case "user":

		fmt.Println(role)
		// defaultRoute := "/user/profile"
		// routes := []string{}
		// roleroutes := newRoleRouter(defaultRoute, routes)
		// c.JSON(
		// 	201,
		// 	gin.H{
		// 		"routes": roleroutes,
		// 	},
		// )

	case "admin":

		fmt.Println(role)
		// defaultRoute := "/admin/profile"
		// routes := []string{}
		// roleroutes := newRoleRouter(defaultRoute, routes)
		// c.JSON(
		// 	201,
		// 	gin.H{
		// 		"routes": roleroutes,
		// 	},
		// )

		// default:
		// 	// TODO: make this return empty values
		//
		// 	c.JSON(
		// 		403,
		// 		gin.H{
		//                "role": role,
		//            },
		// 	)

	}

}
