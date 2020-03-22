package main

import (
	"github.com/roger8b/imc_calc/routers"
	"log"
	"net/http"
)

func main() {

	router := routers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}


