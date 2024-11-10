package controllers

import (
	"net/http"
)

type IController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
