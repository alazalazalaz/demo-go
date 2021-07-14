package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func init(){
	fmt.Println("333333333")
}

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"pong"})
	})


	r.GET("/test", func(c *gin.Context){routeGetTest(c)})
	r.POST("/test", func(c *gin.Context){routePostTest(c)})
	r.Run(":8080")
}

func routeGetTest(c *gin.Context){
	fmt.Println("test")
	name := c.Param("name")
	pw := c.Query("pw")
	nick := c.PostForm("nick")


	fmt.Println(name, pw, nick)
	c.JSON(200, gin.H{"message":"test"})
}

func routePostTest(c *gin.Context){
	fmt.Println("test")
	name := c.Param("name")
	pw := c.Query("pw")
	nick := c.DefaultPostForm("nick", "default")


	fmt.Println(name, pw, nick)
	c.JSON(200, gin.H{"message":"test"})
}


func test(){

}