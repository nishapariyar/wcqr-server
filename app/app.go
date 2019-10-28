package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Api-key", "Content-Type", "Accept"},
	})
	fmt.Println(http.ListenAndServe(":8080", c.Handler(GetRouter())))
}

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/attendees", getAttendees).Methods("GET")
	r.HandleFunc("/attendees", addNewAttendee).Methods("POST")

	r.HandleFunc("/attendees/{id}", getAttendee).Methods("GET")
	r.HandleFunc("/attendees/{id}", updateAttendee).Methods("PUT")

	return r

}
