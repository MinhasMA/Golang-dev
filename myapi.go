package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()                                         //Initializes new router
	r.HandleFunc("/GetInfo", GetInfom).Methods("GET", "OPTIONS") // initializes API
	http.Handle("/", r)                                          // r is like a router. It is a root node - Get info is a child node. "/" is current directory. r is new router. When someone comes to server, it sees this.
	http.ListenAndServe(":8089", nil)                            //http is a package, listenandserve is a function coming from library. This is where you can access it on 8089. This is where
	//I am serving.

}

func GetInfom(w http.ResponseWriter, r *http.Request) {

	//Request is reading. So it needs to be a pointer. ResponseWriter is writing. So its good.

	fmt.Println("Inside GetInfom Function")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		fmt.Println("Inside Get Method")

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
		//http.StatusRequestURITooLong
	}
}
