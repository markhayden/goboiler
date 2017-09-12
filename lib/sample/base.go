package sample

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/markhayden/goboilerplate/lib/config"
	"github.com/markhayden/goboilerplate/lib/models"
	"github.com/markhayden/goboilerplate/lib/util"
)

func InitRouteHandlers(m *pat.PatternServeMux) {
	m.Get("/sample", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SampleHandler(w, r)
	}))
}

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	payload := models.Sample{
		Name:  "Daenerys Targaryen",
		Title: "Mother of Dragons",
		Key:   config.Conf.Keys.B,
	}

	util.WriteJSON(w, r, 200, payload)
	return
}
