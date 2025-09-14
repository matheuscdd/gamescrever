package main

import (
	"fmt"
	"net/http"
	"github.com/matheuscdd/gamescrever/api/storage/databases"
	"github.com/matheuscdd/gamescrever/api/environment"
	"github.com/matheuscdd/gamescrever/api/models/questions"
)

func main() {
	databases.ConnectDatabases()
	env, _ := environment.LoadEnv()
	
	var a []questions.Question
	b := questions.MultipleChoice{}
	b.InsertId()

	a = append(a, &b)
	fmt.Println(b.Id, env.AWSAccessKey)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olág7  5Mund o2163!"))
	})

	http.ListenAndServe(":8000", nil)
}
