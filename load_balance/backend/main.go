package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request handled by %s", hostname)

		fmt.Fprintf(
			w,
			"Backend: %s\nClient: %s\n",
			hostname,
			r.RemoteAddr,
		)
	})

	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)

		fmt.Fprintf(w, "Slow response from %s\n", hostname)
	})

	log.Println("Server", hostname, "running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
