package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Rout = routCtrl{}

type routCtrl struct{}

func (c *routCtrl) GetParametersAccurateInPath(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, "Hello %s", name)
}

func (c *routCtrl) GetParametersFuzzyInPath(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	message := name + " is " + action
	ctx.String(http.StatusOK, message)
}

func (c *routCtrl) GetFullPath(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.FullPath())
}

func (c *routCtrl) GetQueryString(ctx *gin.Context) {
	firstname := ctx.DefaultQuery("firstname", "Guest")
	lastname := ctx.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	ctx.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func (c *routCtrl) PostForm(ctx *gin.Context) {
	id := ctx.Query("id")

	page := ctx.DefaultQuery("page", "0")
	name := ctx.PostForm("name")
	message := ctx.PostForm("message")

	ctx.JSON(200, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func (c *routCtrl) PostMap(ctx *gin.Context) {
	ids := ctx.QueryMap("ids")
	names := ctx.PostFormMap("names")

	ctx.String(http.StatusOK, "ids: %v; names: %v", ids, names)
}

func (c *routCtrl) UploadPicture(ctx *gin.Context) {
	// single file
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	ctx.SaveUploadedFile(file, "/Users/chenpeng/working/own/pic/"+strconv.FormatInt(time.Now().Unix(), 10)+file.Filename)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (c *routCtrl) UploadPictures(ctx *gin.Context) {
	// Multipart form
	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	for _, file := range files {
		log.Println(file.Filename)

		filePath := "/Users/chenpeng/working/own/pic/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
		// Upload the file to specific dst.
		ctx.SaveUploadedFile(file, filePath)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
