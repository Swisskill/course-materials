package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

var LOG_LEVEL int = 2

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	var fCount int
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")
		//for each regex, print it out
		//furthermore, store it
		for _, r := range regexes {
			if r.MatchString(path) {
				var tfile FileInfo
				dir, filename := filepath.Split(path)
				tfile.Filename = string(filename)
				tfile.Location = string(dir)

				//TODO_5: As it currently stands the same file can be added to the array more than once
				//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
				flag := false

				for i := range Files {
					if tfile == Files[i] {
						flag = true
					}
				}
				if !flag {

					Files = append(Files, tfile) //this is number 5. this line is what appends the tfile to the Files
					fCount++
				}
				if w != nil && len(Files) > 0 {

					//TODO_6: The current key value is the LEN of Files (this terrible);
					//TODO_6: Create some variable to track how many files have been added
					w.Write([]byte(`"` + (strconv.FormatInt(int64(fCount), 10)) + `":  `))
					json.NewEncoder(w).Encode(tfile)
					w.Write([]byte(`,`))

				}
				if LOG_LEVEL > 0 {
					log.Printf("[+] HIT: %s\n", path)
				}
			}

		}
		return nil
	}

}

//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function,
//TODO_7: Instead of using the regexes array, define a single regex
//TODO_7: Hint look at the logic in scrape.go to see how to do that;
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		r := regexp.MustCompile(`(?i` + query)
		if r.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)

			if w != nil && len(Files) > 0 {

				//TODO_6: The current key value is the LEN of Files (this terrible);
				//TODO_6: Create some variable to track how many files have been added
				w.Write([]byte(`"` + (strconv.FormatInt(int64(len(Files)), 10)) + `":  `))
				json.NewEncoder(w).Encode(tfile)
				w.Write([]byte(`,`))

			}
			if LOG_LEVEL > 0 {
				log.Printf("[+] HIT: %s\n", path)
			}
		}
		return nil
	}
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status" : "API is up and running ",`))
	var regexstrings []string

	for _, regex := range regexes {
		regexstrings = append(regexstrings, regex.String())
	}

	w.Write([]byte(` "regexs" :`))
	json.NewEncoder(w).Encode(regexstrings)
	w.Write([]byte(`}`))
	if LOG_LEVEL > 0 {
		log.Println(regexes)
	}

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	//TODO_8 - Write out something better than this that describes what this api does

	fmt.Fprintf(w, htm)
}

func FindFile(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	q, ok := r.URL.Query()["q"]

	w.WriteHeader(http.StatusOK)
	if ok && len(q[0]) > 0 {
		if LOG_LEVEL > 0 {
			log.Printf("Entering search with query=%s", q[0])
		}
		// ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
		// e.g., func finder(query string) []string { ... }

		for _, File := range Files {
			if File.Filename == q[0] {
				json.NewEncoder(w).Encode(File.Location)
				//consider FOUND = TRUE
			}
		}
		//TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...

	} else {
		// didn't pass in a search term, show all that you've found
		w.Write([]byte(`"files":`))
		json.NewEncoder(w).Encode(Files)
	}
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")

	location, locOK := r.URL.Query()["location"]
	reger, regerOk := r.URL.Query()["regex"]
	//TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
	//TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
	//TODO_10: something like ...  rootDir string := "???"
	//TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole

	if locOK && len(location[0]) > 0 {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{ "parameters" : {"required": "location",`))
		w.Write([]byte(`"optional": "regex"},`))
		w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
		w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
		return
	}

	//wrapper to make "nice json"
	w.Write([]byte(`{ `))

	// TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
	// Define the logic required here to call the new function walkFn2(w,regex[0])
	// Hint, you need to grab the regex parameter (see how it's done for location above...)
	//11 requires 7

	// if regexOK
	//   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
	// else run code to locate files matching stored regular expression
	//if err := filepath.Walk(location[0], walkFn(w)); err != nil { //w happens to be the http response writer
	//this is also where todo 10 would be so that you can't search from the root

	//	log.Panicln(err)
	//}
	direct := `/Users/wrbra/Desktop/COSC/Cyber`
	if regerOk {
		//log.Printf("walkfn2")
		filepath.Walk(direct+location[0], walkFn2(w, reger[0]))
	} else {
		//log.Printf("walk")
		filepath.Walk(direct+location[0], walkFn(w))
	}

	//wrapper to make "nice json"
	w.Write([]byte(` "status": "completed"} `))

}

//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!
//todo12
func ResetRex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	resetRegEx()
	Files = nil
}

//todo13
func ClearRex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	clearRegEx()
}

//todo14
func AddRex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL > 0 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//params["regex"]
	addRegEx(`(?i)` + params["regex"]) //dr mike recommends better error checking here
}

//========================================================================//
//TODO_13 create endpoint that calls clearRegEx ;
//TODO_12 Make sure to connect the name of your function back to the clear endpoint main.go!

//TODO_14 create endpoint that calls addRegEx ;
//TODO_12 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint
const htm = `<html>    
<head>      
   <title>Hacking Your Computer</title>    
</head>    
<body style="background-color:black;">      
   <h1 style="font-family:Consolas;color:red;font-size:40px;">       Welcome to the page devoted to hacking yourself!</h1>      
   <p style="font-family:Courier New; color:blue;font-size:18px;">Hacking your computer is as simple as these few commands.
   Basically, the following instructions will allow you view your own files. The commands available are: </p> 
   <p style="font-family:Courier New; color:white;font-size:20px;">/api-status </p>
   <p style="font-family:Courier New; color:white;font-size:20px;">/indexer </p>
   <p style="font-family:Courier New; color:white;font-size:20px;">/search </p>
   <p style="font-family:Courier New; color:white;font-size:20px;">/addsearch/{regex} </p>
   <p style="font-family:Courier New; color:white;font-size:20px;">/clear </p>
   <p style="font-family:Courier New; color:white;font-size:20px;">/reset </p>
      
</body>
</html>`
