package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kajikentaro/playground/str-to-img/controllers"
	"github.com/kajikentaro/playground/str-to-img/services"
)

func main() {
	service := services.NewGenImageService()
	controller := controllers.NewGenImageController(service)

	r := mux.NewRouter()
	r.HandleFunc("/", controller.GenImageHandler).Methods(http.MethodGet)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
