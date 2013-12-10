package main

import (
	"github.com/codegangsta/martini"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "hello world"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Get("/hello", Auth, func(params martini.Params) string {
		return "Hello random person"
	})

	m.Get("/auth", func() string {
		return "hello world"
	})

	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "secret" {
		http.Error(res, "Nope nope nope", 401)
	}
	res.Header().Add("API-REPLY", "OK")
}
