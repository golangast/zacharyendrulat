package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("works")

	log.Fatal(http.ListenAndServe(":8081", ))
}
