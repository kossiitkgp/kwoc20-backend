package routes

import (
	"github.com/gorilla/mux"

	// TEMP
	"kwoc20-backend/controllers"
	"kwoc20-backend/utils"
)

// TEMP
// Discuss and add 2 middlewares - JWT Required, JSON Marshalling
func RegisterMentor(r *mux.Router) {
	r.HandleFunc("/form", utils.PermissiveCORSMiddleware(utils.LoginRequired(utils.JsonIO(controllers.MentorReg)))).Methods("POST")
	r.HandleFunc("/dashboard", utils.PermissiveCORSMiddleware(utils.JsonIO(controllers.MentorDashboard))).Methods("POST")

}
