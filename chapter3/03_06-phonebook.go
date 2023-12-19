package main

import (
	"os"
	"text/template"
)

type contact struct {
	Name        string
	PhoneNumber string
	Email       string
	Country     string
	Category    string
}

func main() {
	// Create a new phone book
	book := []contact{
		{
			Name:        "John Doe",
			PhoneNumber: "1234567890",
			Email:       "johnd@jdoe.com",
			Country:     "USA",
			Category:    "Family",
		},
		{
			Name:        "Jane Doe",
			PhoneNumber: "1234567890",
			Email:       "janed@jdoe.com",
			Country:     "USA",
			Category:    "Work",
		},
		{
			Name:        "Juan Doe",
			PhoneNumber: "1234567890",
			Email:       "juand@jdoe.com",
			Country:     "India",
			Category:    "Family",
		},
	}

	var templFile = "03_06-phonebook.tmpl"

	tmpl, err := template.New(templFile).ParseFiles(templFile)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, book)
	if err != nil {
		panic(err)
	}
}
