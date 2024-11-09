package routes

import "github.com/gorilla/mux"

type IRoutes interface {
	RegisterRoutes(r *mux.Router)
}
