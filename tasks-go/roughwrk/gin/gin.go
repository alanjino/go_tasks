package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
) /* func CheckOdd(c chan int, d *gin.Context) {
    val := <-c
    switch {
//  case val == 999999:
//      fmt.Println("favicon request")
    case val%2 == 0:
        fmt.Println("its a even number")
        d.JSON(200, gin.H{
            "response": "its a even number",
        })
    case val%2 != 0:
        fmt.Println("its a odd number")
        d.JSON(200, gin.H{
            "response": "its a odd number",
        })    default:
*/ //fmt.Println("")
// }
// }
func getParam(c *gin.Context) {
	channel := make(chan int)
	defer close(channel) //  go CheckOdd(channel, c)
	chekodd(channel)
	id := c.Param("value")
	if id == "favicon.ico" {
		channel <- 999999
	} else {
		val, _ := strconv.Atoi(id)
		channel <- val

	}
}
func chekodd(c chan int) {
	val := <-c
	switch {
	case val%2 == 0:
		fmt.Println("its a even number")
	case val%2 != 0:
		fmt.Println("its a odd number")
	default:
		fmt.Println("")
	}
}
func main() {
	r := gin.Default()
	r.GET("/value", getParam)
	r.Run(":4000")
}
