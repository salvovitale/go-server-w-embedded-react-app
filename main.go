package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed ui/build
	web embed.FS
)

func main() {

	dist, _ := fs.Sub(web, "ui/build")
	http.Handle("/", http.FileServer(http.FS(dist)))

	fmt.Println("Your embedded react app is running on http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))

}
