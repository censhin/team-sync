package main

import (
    "github.com/gin-gonic/gin"
    ctrl "./controllers"
)

func main() {
    r := gin.Default()

    r.GET("/", ctrl.Hello)
    r.GET("/users", ctrl.GetUsers)

    r.Run(":9000")
}