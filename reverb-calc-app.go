package main

import (
	"controller"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const defaultServerPort = 9090

// Params
var serverPort int

func main() {
	initFlags()
	serverListenConfig := fmt.Sprintf(":%d", serverPort)

	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/reverb-calc", controller.Index)
	r.PathPrefix("/reverb-calc/css/").Handler(http.StripPrefix("/reverb-calc/css/", http.FileServer(http.Dir("src/view/assets/css"))))
	r.PathPrefix("/reverb-calc/js/").Handler(http.StripPrefix("/reverb-calc/js/", http.FileServer(http.Dir("src/view/assets/js"))))
	http.Handle("/", r)

	fmt.Printf("starting Server at %v\n", serverListenConfig)
	err := http.ListenAndServe(serverListenConfig, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initFlags() {
	flag.IntVar(&serverPort, "port", defaultServerPort, "The Port the server should listen.")
	flag.Parse()
}
