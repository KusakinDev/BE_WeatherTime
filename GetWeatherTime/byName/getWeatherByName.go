package byname

import (
	"encoding/json"
	"log"
	owmstr "main/Struct/owmStr"
	"main/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherByName(c *gin.Context) {

	city := c.Query("name")
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

	var weatherData owmstr.WeatherData
	err = json.NewDecoder(response.Body).Decode(&weatherData)
	if err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %s", err)
	}

	log.Println(weatherData)

}
