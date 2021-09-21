package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context){
	name := c.Query("name")
	str := fmt.Sprintf("hello i am %s ! \n",name)
	fmt.Println(str)

	c.String(http.StatusOK,str)

}