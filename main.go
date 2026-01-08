// @title User API
// @version 1.0
// @description Simple CRUD User tanpa DB (In-Memory)
// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"net/http"

	"github.com/FikriBaihaqi73/go-study/routes"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/FikriBaihaqi73/go-study/docs"
)

func main() {
	r := routes.SetupRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Server jalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
