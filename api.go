package main

import (
	"log"
	byname "main/GetWeatherTime/byName"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/weather", byname.GetWeatherByName)

	log.Println("Server starting at :8080")
	log.Fatal(r.Run(":8080"))
}
