package controller

import "net/http"

type Controller struct {
	ControllerMethod
	Path string
}

type ControllerMethod interface {
	Handle(w http.ResponseWriter, r *http.Request)
}
