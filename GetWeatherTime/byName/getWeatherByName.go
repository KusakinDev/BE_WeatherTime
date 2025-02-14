package byname

import (
	"log"
	"main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherByName(c *gin.Context) {

	city := c.DefaultQuery("city", "Unknown")
	log.Println(city)

	if city == "" {
		c.JSON(400, gin.H{"error": "Name param is missing"})
		return
	}

	apiURL := config.WEATHER_API_1 + city + config.WEATHER_API_2

	log.Println(apiURL)

	response, err := http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WatherAPI error"})
		return
	}
	defer response.Body.Close()

	response, err = http.Get(apiURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WatherAPI error"})
		return
	}
	defer response.Body.Close()

	log.Println(response.Body)

}
