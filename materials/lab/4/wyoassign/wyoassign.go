package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"

	//"main"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	//"main/main"
	"github.com/gorilla/mux"
)

type Response struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id          int    `json:"id"`
	Title       string `json:"title`
	Description string `json:"desc"`
	Points      int    `json:"points"`
}

var Assignments []Assignment

const Valkey string = "FooKey"

func InitAssignments() {
	//----We do a little random number generation
	rand.Seed(time.Now().UnixNano())
	ranId := rand.Intn(1024)
	var assignmnet Assignment
	assignmnet.Id = ranId
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	Assignments = append(Assignments, assignmnet)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		paramI, _ := strconv.Atoi(params["id"])
		if assignment.Id == paramI {
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	//w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	//if main.Pass != "mikeiscool" {
	//	return
	//} //this should implement a password function. Hopefully
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	if r.FormValue("password") != "mikeiscool" {
		log.Printf("You have not entered the correct password")
		return
	}

	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
		paramI, _ := strconv.Atoi(params["id"])
		if assignment.Id == paramI {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	//got some help on this part from Aram. Not sure if it actually works right though but I have to go now
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")

	var response Response
	response.Assignments = Assignments
	r.ParseForm()

	if r.FormValue("password") != "mikeiscool" {
		log.Printf("FAIL: incorrect or missing key!")
		w.WriteHeader(http.StatusBadRequest)
		response := make(map[string]string)
		response["status"] = "Wrong or missing key value"
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return
		}
		w.Write(jsonResponse)
		return
	}

	DeleteAssignment(w, r)
	var updated Assignment
	if r.FormValue("id") != "" {
		updated.Id, _ = strconv.Atoi(r.FormValue("id"))
		updated.Title = r.FormValue("title")
		updated.Description = r.FormValue("desc")
		updated.Points, _ = strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, updated)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, assignment := range Assignments {
		check, _ := strconv.Atoi(r.FormValue("id"))
		if assignment.Id == check {
			log.Printf("Sorry. Choose a different id please")
			w.WriteHeader(http.StatusNotFound)
			//json.NewEncoder(w).Encode(assignment)
			return
		}
	}
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if r.FormValue("id") != "" {

		//forming, _ := strconv.Atoi(r.FormValue("id")) //this line is stupid. Maybe it means the lines above don't work idk
		assignmnet.Id, _ = strconv.Atoi(r.FormValue("id"))
		assignmnet.Title = r.FormValue("title")
		assignmnet.Description = r.FormValue("desc")
		assignmnet.Points, _ = strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}
