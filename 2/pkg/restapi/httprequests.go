package restapi

import "fmt"
import "log"
import "net/http"
import "github.com/gorilla/mux"

func StartRouter() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeLink)
	router.HandleFunc("/status", StatsPage)
	router.HandleFunc("/nom", NomPage)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Another page said Welcome home!")
}

func StatsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the status page")
}

func NomPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the nomnom page")
}