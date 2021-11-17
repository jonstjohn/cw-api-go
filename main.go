// user=cwapp password=%s database=cw host=/cloudsql/api-project-XXXX:us-central1:cw-pg-dev # /.s.PGSQL.5432

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	r := gin.Default()
	r.GET("/area/:areaId/forecast", func(c *gin.Context) {

		areaId, err := strconv.Atoi(c.Param("areaId"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Area ID is not an integer: %v\n", err)
			os.Exit(1)
		}
		url := os.Getenv("DB_URL")

		conn, err := pgx.Connect(context.Background(), url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer conn.Close(context.Background())

		var name string
		var latitude float32
		var longitude float32
		err = conn.QueryRow(context.Background(), "select name, latitude, longitude from area where area_id=$1", areaId).Scan(&name, &latitude, &longitude)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		/*
			rows, err := conn.Query(context.Background(),
				"SELECT date, high, low, precip_day/100, precip_night/100, sky/100, relative_humidity/100,"+
					" wsym, rain_amount, snow_amount, wind_sustained, wind_gust, weather"+
					" FROM daily WHERE area_id = $1 ORDER BY date asc", areaId)
			if err != nil {
				fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			defer rows.Close()

			daily := make(map[string]interface{})
			daily["summary"] = ""
			data := make([]map[string]interface{}, 0)
			for rows.Next() {
				var date time.Time
				var high int
				var low int
				var precipDay float32
				var precipNight float32
				var sky float32
				var humidity float32
				var wsym string
				var rainAmount float32
				var snowAmount float32
				var windSustained int
				var windGust int
				var weather string
				//values, err := rows.Values()
				rows.Scan(&date, &high, &low, &precipDay, &precipNight, &sky, &humidity, &wsym, &rainAmount, &snowAmount, &windSustained, &windGust, &weather)
				day := make(map[string]interface{})
				day["time"] = date.Unix()
				day["temperatureHigh"] = high
				day["temperatureLow"] = low
				day["precipProbability"] = precipDay
				day["precipProbabilityNight"] = precipNight
				day["humidity"] = humidity
				day["windSpeed"] = windSustained
				day["windGust"] = windGust
				data = append(data, day)
			}
		*/

		/*
			daily["summary"] = ""
			data := make([]map[string]interface{}, 0)
			day0 := make(map[string]interface{})
			day0["time"] = 1636970400
			day0["temperatureHigh"] = 72
			day0["temperatureLow"] = 51
			day0["precipProbability"] = 0.18
			day0["precipProbabilityNight"] = 0.20
			day0["humidity"] = 0.72
			day0["windSpeed"] = 2
			day0["windGust"] = 9
			data = append(data, day0)
		*/
		//daily["data"] = data

		c.JSON(200, gin.H{
			"name":      name,
			"latitude":  latitude,
			"longitude": longitude,
			//"daily":     daily,
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("0.0.0.0:%v", port)

	r.Run(address) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
