package main

import (
	"github.com/gin-gonic/gin"
	"gopractice/projects/gindemo/control"
	"net/http"
)

func main() {
	router := gin.Default()

	router.OPTIONS("/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "all")
	})

	parametersInPathGroup := router.Group("/user")
	{
		// This handler will match /user/john but will not match /user/ or /user
		parametersInPathGroup.GET("/:name", control.Rout.GetParametersAccurateInPath)

		// However, this one will match /user/john/ and also /user/john/send
		// If no other routers match /user/john, it will redirect to /user/john/
		parametersInPathGroup.GET("/:name/*action", control.Rout.GetParametersAccurateInPath)

		// For each matched request Context will hold the route definition
		parametersInPathGroup.POST("/:name/*action", control.Rout.GetFullPath)
	}

	routerGroup := router.Group("")
	{
		// Query string parameters are parsed using the existing underlying request object.
		// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
		routerGroup.GET("/welcome", control.Rout.GetQueryString)

		//POST /post?id=1234&page=1 HTTP/1.1
		//Content-Type: application/x-www-form-urlencoded
		//
		//name=manu&message=this_is_great
		routerGroup.POST("/post/form", control.Rout.PostForm)

		//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
		//Content-Type: application/x-www-form-urlencoded
		//
		//names[first]=thinkerou&names[second]=tianou
		routerGroup.POST("/post/map", control.Rout.PostMap)

		routerGroup.POST("/upload/picture", control.Rout.UploadPicture)
		routerGroup.POST("/upload/pictures", control.Rout.UploadPictures)
	}

	servingGroup := router.Group("")
	{
		servingGroup.GET("/local/file", func(c *gin.Context) {
			c.File("go.mod")
		})
	}

	_ = router.Run(":10000")
}
