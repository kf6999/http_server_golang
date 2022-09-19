package main

import (
	"fmt"
	database "github.com/kf6999/http_server_golang/internal"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("server started on ", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

	c := database.NewClient("db.json")
	err = c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
