package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("./templates/cunningham.tmpl")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "HOME",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}