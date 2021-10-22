package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"path"
)

func init(){
	fmt.Println("333333333")
}

func main(){
	r := gin.New()
	//r.GET("/ping", func(c *gin.Context){
	//	c.JSON(200, gin.H{"message":"pong"})
	//})


	r.GET("/user/:name", routeUser)
	r.POST("/test", routePostTest)
	r.Run(":8080")
}

func routeUser(c *gin.Context){
	fmt.Println("test")
	name := c.Param("name")
	pw := c.Query("pw")
	nick := c.PostForm("nick")

	path1, path2, path3 := "/aaa/", "bbb", "cc"
	log.Printf("%s\n", path.Join(path1, path2, path3))

	fmt.Println(name, pw, nick)
	c.JSON(200, gin.H{"message":name})
}

func routeAge(c *gin.Context){
	fmt.Println("test")
	age := c.Param("age")
	pw := c.Query("pw")
	nick := c.PostForm("nick")

	path1, path2, path3 := "/aaa/", "bbb", "cc"
	log.Printf("%s\n", path.Join(path1, path2, path3))

	fmt.Println(pw, nick)
	c.JSON(200, gin.H{"message":age})
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