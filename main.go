package main

import (
	"fmt"
	"log"
	"net/http"
	"wordscloud/root"
)

func main() {
	http.HandleFunc("/", root.RootHandle)

	log.Fatal(http.ListenAndServe(":50000", nil))
	fmt.Println("Started listening to the port 50000")
}
