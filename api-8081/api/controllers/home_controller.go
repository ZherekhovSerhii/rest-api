package controllers

import (
	"net/http"

	"github.com/ZherekhovSerhii/http-rest-api/api/responses"
)

// Home ...
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
