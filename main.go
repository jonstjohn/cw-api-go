// Sample run-helloworld is a minimal Cloud Run service.
//package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	log.Print("starting server...")
// 	http.HandleFunc("/", handler)

// 	// Determine port for HTTP service.
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 		log.Printf("defaulting to port %s", port)
// 	}

// 	// Start HTTP server.
// 	log.Printf("listening on port %s", port)
// 	if err := http.ListenAndServe(":"+port, nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	name := os.Getenv("NAME")
// 	if name == "" {
// 		name = "World"
// 	}
// 	fmt.Fprintf(w, "Hello %s!\n", name)
// }

package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("0.0.0.0:%v", port)

	r.Run(address) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
