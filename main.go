// user=cwapp password=%s database=cw host=/cloudsql/api-project-736062072361:us-central1:cw-pg-dev # /.s.PGSQL.5432

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
		err = conn.QueryRow(context.Background(), "select name from area where area_id=$1", areaId).Scan(&name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": name,
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("0.0.0.0:%v", port)

	r.Run(address) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
