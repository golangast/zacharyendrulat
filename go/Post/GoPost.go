package GoPost

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func UnmarshalLogin(data []byte) (Postcms, error) {
	var r Postcms
	fmt.Print("starting unmarshal")
	err := json.Unmarshal(data, &r)
	fmt.Print("is starting", data)

	return r, err
}

func (r *Postcms) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Postcms struct {
	Date  string `json:"date"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Html  string `json:"html"`
}

func GoPosts(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081/post")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

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
		w.Write([]byte("Received a GET request\n"))
		fmt.Println("reqeusted boyd ", r.Body)
		w.WriteHeader(http.StatusOK)

	case "POST":
		fmt.Println(r.Header.Get("Origin"))

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
		fmt.Println("reqeusted boyd ", r.Body)
		l, err := UnmarshalLogin(reqBody)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("opening database")

		//database beginsssssss

		db, err := sql.Open("mysql", "root:@/users")
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
		// query
		stmt, err := db.Prepare("INSERT INTO post(date, title, slug, html) VALUES(?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		userstemp := Postcms{Date: l.Date, Title: l.Title, Slug: l.Slug, Html: l.Html}
		fmt.Println(userstemp)
		res, err := stmt.Exec(userstemp.Date, userstemp.Title, userstemp.Slug, userstemp.Html)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println("reached query")

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
