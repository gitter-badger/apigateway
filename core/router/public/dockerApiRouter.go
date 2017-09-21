package public

import (
	"apigateway/core/api/docker"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"log"
	"net/http"
)

func PublicApiRouter(router gin.Engine) {
	log.Println("start init public router.......")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	router.GET("/getImageTagInfo", func(c *gin.Context) {
		log.Println(">>>>>>>" + c.DefaultQuery("imageTag", "test"))
		imageName := c.DefaultQuery("imageType", "") + "/" + c.DefaultQuery("imageName", "")
		imageTag := c.DefaultQuery("imageTag", "")
		str := docker.GetImageTagInfo(imageName, imageTag)
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "imageTagList": string(str)})
	})

	router.GET("/getImageAllTags", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		imageName := c.DefaultQuery("imageName", "")
		str := docker.GetAllTagByImageName(imageName)
		c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "imageTagList": string(str)})

	})

	router.POST("/form_post", func(c *gin.Context) {
		log.Println(c.Request.Header)
		log.Println(c.Request.Body)
		htmlData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Println(err)

		}
		// print out
		log.Println(string(htmlData)) //<-- here !
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		log.Println(">>>>>>>>>>>" + message)
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	log.Println("complete init public router.......")
}
