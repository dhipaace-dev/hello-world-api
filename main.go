package main

import (
  "os"
  "hello-world-api/controllers"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/hello", controllers.HelloWorld)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router.Run(":" + port)
}


