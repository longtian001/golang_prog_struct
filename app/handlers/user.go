package handlers

import (
	"app/models"
	"net/http"
)

func AllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	users, err := models.AllUsers()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	JsonResponse(users, w)
}
