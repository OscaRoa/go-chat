package main

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, r := range routes{
		router.Handle(r.Method, r.Pattern, r.HandlerFunc)
	}

	return router
}
