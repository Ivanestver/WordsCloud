package main

import (
	"fmt"
	"log"
	"net/http"
	"wordscloud/poll"
	"wordscloud/root"
)

func main() {
	fs := http.FileServer(http.Dir("poll"))
	http.Handle("/poll/", http.StripPrefix("/poll/", fs))
	http.HandleFunc("/root", root.RootHandle)
	http.HandleFunc("/new_poll", poll.HandleNewPoll)
	http.HandleFunc("/create_poll", poll.HandleCreatePoll)

	log.Fatal(http.ListenAndServe(":50000", nil))
	fmt.Println("Started listening to the port 50000")
}
