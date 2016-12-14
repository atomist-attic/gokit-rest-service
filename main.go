package main

import (
	"log"
	"net/http"

	"github.com/atomist-project-templates/gokit-rest-service/mymicroservice"
)

func main() {
	mymicroservice.AddServices()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
