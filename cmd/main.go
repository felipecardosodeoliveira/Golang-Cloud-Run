package main

import (
	"fmt"
	"github/felipecardosodeoliveira/golang-cloud-run/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cep/", handler.CepHandler)
	port := "8080"
	fmt.Printf("Server is running %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
