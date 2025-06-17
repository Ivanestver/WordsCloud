package main

import (
	"fmt"
	"log"
	"net/http"
	"wordscloud/poll"
	"wordscloud/root"
)

func main() {
	http.HandleFunc("/", root.RootHandle)
	http.HandleFunc("new_poll", poll.HandleNewPoll)

	log.Fatal(http.ListenAndServe(":50000", nil))
	fmt.Println("Started listening to the port 50000")
}
