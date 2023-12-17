// Chipi chipi chapa chapa
// Dubi dubi daba daba
// Magico mi dubi dubi
// Boom boom boom boom
package main

import (
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robots.txt")
	})

    http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "video.mp4")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    http.ListenAndServe(":"+port, nil)
}
