package main

import (
	"fmt"
	"github.com/matheuscdd/gamescrever/api/models/questions"
)



func main() {
	var a []questions.Question
	b := questions.MultipleChoice{}
	b.InsertId()

	a = append(a, &b)
	fmt.Println(b.Id)
}