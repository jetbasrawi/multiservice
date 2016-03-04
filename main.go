package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

type PageData struct {
	NumViews string
}

func main() {

	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Set up connection to redis
		conn, err := redis.Dial("tcp", "redis:6379")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//Close the connection when completed
		defer conn.Close()

		//Increment the number of views
		conn.Do("INCR", "numviews")

		//Get the current number of views
		numviews, err := redis.String(conn.Do("GET", "numviews"))

		pageData := PageData{NumViews: numviews}

		if err := templates.ExecuteTemplate(w, "index.html", pageData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("Numviews " + numviews)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
