package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Can't read hostname")
	}

	m := martini.Classic()
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(res, "Hello from Go!<br>Server listening port %s on host: %s", port, hostname)
	})
	m.RunOnAddr(":" + port)
}
