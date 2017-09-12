package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/markhayden/goboilerplate/lib/config"
	"github.com/markhayden/goboilerplate/lib/models"
	"github.com/markhayden/goboilerplate/lib/sample"
	"github.com/markhayden/goboilerplate/lib/util"
)

func main() {
	// initialize the pat muxer
	m := pat.New()

	// handle requests to the root as an example
	m.Get("/", http.HandlerFunc(rootHandler))

	// initialize additional handlers (example)
	sample.InitRouteHandlers(m)

	// start the api server
	log.Print("api server running at localhost:12345")
	http.Handle("/", m)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// rootHandler handles the output of requests made to the root domain
func rootHandler(w http.ResponseWriter, r *http.Request) {
	payload := models.Sample{
		Name:  "Jon Snow",
		Title: "King of the North",
		Key:   config.Conf.Keys.A,
	}

	util.WriteJSON(w, r, 200, payload)
	return
}
