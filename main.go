package main

import (
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
)

func handler(c *gin.Context) {
	body, _ := c.Request.GetBody()
	log.Printf("body == %#v", body)
	c.JSON(200, "")
}
func main() {
	router := gin.Default()
	router.GET("/", handler)
	gateway.ListenAndServe("", router)
}
