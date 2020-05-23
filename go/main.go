package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	Con "github.com/golangast/zacharyendrulat/go/Context"
	GET "github.com/golangast/zacharyendrulat/go/GET"
	P "github.com/golangast/zacharyendrulat/go/Post"
	"github.com/rs/cors"
)

var err error

func main() {

	fmt.Println("starting server......")
	mux := http.NewServeMux()
	mux.HandleFunc("/post", P.GoPosts)
	mux.HandleFunc("/get", GET.GoGET)

	//used to get other files css/js
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//used for keeping context and requests for the server for logging
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Con.AddContext(c, handler)))

}
