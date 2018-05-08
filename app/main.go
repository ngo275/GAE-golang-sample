package main

import (
	"net/http"

	"github.com/ngo275/gae-blueprint/src/handler"
	"google.golang.org/appengine"
)

func init() {
	r := handler.Router()
	http.Handle("/", r)

	appengine.Main()
}
