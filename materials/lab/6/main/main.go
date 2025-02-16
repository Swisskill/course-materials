package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"scrape/scrape"

	"github.com/gorilla/mux"
)

var LOG_LEVEL int = 2

//LOG_LEVEL := 2
//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)
//ATTN Dr. Mike: I was told that the log level thing doesn't actually do anything for us, but that it's just for future configuration
//If this was wrong, well dang, I did it wrong. Hopefully it's not that

func main() {

	if LOG_LEVEL > 0 {
		log.Println("starting API server")
	} //these are the logs mentioned in todo 1.
	//log level. set it global. change it as needed
	//more realistic would pass in some variable from the environment
	//create a function you pass the logs to to check via if what log level it is
	//create a new router
	router := mux.NewRouter()
	if LOG_LEVEL > 0 {
		log.Println("creating routes")
	}
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	router.HandleFunc("/addsearch/{regex}", scrape.AddRex).Methods("GET")
	router.HandleFunc("/clear", scrape.ClearRex).Methods("GET")
	router.HandleFunc("/reset", scrape.ResetRex).Methods("GET")
	//these todos will be super easy once you fix the other todos
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

//C:\Users\wrbra\Desktop\COSC\Cyber\course-materials\materials\lab\6\main\main.go
