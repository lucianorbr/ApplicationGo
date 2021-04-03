package main

import (
	"net/http"
	"os"

	"github.com/gansta/negrori"
	"github.com/gorilla/mux"

)


// NewServer Configures and returns a Server
func NewServer()  *negrori.Negrori {
	n:= negrori.Classic()
	mx := mux.NewRouter()

	initRoutes(mx)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mx.Router, formatter *render.Render)  {
	weRoot = os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		root, err := os.Getwd()
		if err != nil {
			panic("Could not retrieve working diretory")
		} else {
			webRoot = root
		}
	}

	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
