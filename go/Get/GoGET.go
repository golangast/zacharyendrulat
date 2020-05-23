package GoPost

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID    int    `json:"ID"`
	Date  string `json:"Dates"`
	Title string `json:"Title"`
	Slug  string `json:"Slug"`
	Html  string `json:"Html"`
}

func GoGET(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081/post")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/show")
	w.Header().Set("Access-Control-Allow-Origin", "http://1299991b.ngrok.io/get")

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		fmt.Println(r.Header.Get("Origin"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}

		fmt.Println("reqeusted boyd ", r.Body)
		w.WriteHeader(http.StatusOK)

		fmt.Println("opening database")

		//database beginsssssss

		db, err := sql.Open("mysql", "root:@/phpmyadmin")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("open ")
		}

		defer db.Close()
		err = db.Ping()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ping ")
		}

		var (
			id    int
			date  string
			title string
			slug  string
			html  string
		)
		i := 0
		var post []Post

		// query
		rows, err := db.Query("select * from phpmyadmin.post")
		for rows.Next() {
			err := rows.Scan(&id, &date, &title, &slug, &html)
			if err != nil {
				fmt.Println(err)
			} else {
				i++
				fmt.Println("scan ", i)
			}

			post = append(post, Post{ID: id, Date: date, Title: title, Slug: slug, Html: html})
			fmt.Println("before marshal ", post)

		}
		json.NewEncoder(w).Encode(post) //remember to encode it

		defer rows.Close()
		w.Header().Set("Content-Type", "application/json")

		fmt.Println("reached query")

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
