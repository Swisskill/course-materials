package main

import (
	"log"
	"net/http"
	"wyoassign/wyoassign"

	"github.com/gorilla/mux"
)

func main() {
	wyoassign.InitAssignments()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")
	router.HandleFunc("/assignment", wyoassign.CreateAssignment).Methods("POST")
	router.HandleFunc("/assignment", wyoassign.UpdateAssignment).Methods("PUT")
	// router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

	//Now we begin the TUI
	/*
		var pass string
		var useR string
		fmt.Println("(a little note, the password is hardcoded in goofy slice. Just look through the file. Pretend it's in another location)")
		fmt.Println("Please enter your username to continue\n")
		fmt.Scanln(&useR)
		if useR == "drmike" {
			fmt.Println("---Welcome Dr Mike---\nPlease enter your password")
		} else {
			fmt.Println("Please enter your password")
		}
	*/

}
