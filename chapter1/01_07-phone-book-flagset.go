package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
)

const phoneBookPath = "phonebook.json"

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Records struct {
	Entries []Person
}

func (r *Records) AddEntry(record Person) {
	r.Entries = append(r.Entries, record)
}

func (r *Records) DeleteEntry(name string) {
	for i, e := range r.Entries {
		if e.Name == name {
			r.Entries = append(r.Entries[:i], r.Entries[i+1:]...)
			break
		}
	}
}

func (r *Records) GetEntry(name string) []Person {
	if name == "" {
		return r.Entries
	}
	for _, e := range r.Entries {
		if e.Name == name {
			return []Person{e}
		}
	}
	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <subcommand> <options>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Subcommands:\n")
	fmt.Fprintf(os.Stderr, "  add - Add a new phone book record\n")
	fmt.Fprintf(os.Stderr, "     Options:\n")
	fmt.Fprintf(os.Stderr, "       --name - Name of the person\n")
	fmt.Fprintf(os.Stderr, "       --phone - Phone number of the person\n\n")

	fmt.Fprintf(os.Stderr, "  delete - Delete an existing phone book record\n")
	fmt.Fprintf(os.Stderr, "     Options:\n")
	fmt.Fprintf(os.Stderr, "       --name - Name of the person\n")
	fmt.Fprintf(os.Stderr, "       --phone - Phone number of the person\n\n")

	fmt.Fprintf(os.Stderr, "  get - Get one or all phonebook records\n")
	fmt.Fprintf(os.Stderr, "     Options:\n")
	fmt.Fprintf(os.Stderr, "       --name - Name of the person\n")
	flag.PrintDefaults()
}

func main() {
	phoneBook, err := ioutil.ReadFile(phoneBookPath)
	if err != nil {
		if os.IsNotExist(err) {
			phoneBook = []byte("[]")
			ioutil.WriteFile(phoneBookPath, phoneBook, 0644)
		} else {
			panic(err)
		}
	}

	var records Records

	json.Unmarshal(phoneBook, &records.Entries)

	// Define command-line flags
	var name string
	var phone string

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addCmd.StringVar(&name, "name", "", "Name of the person")
	addCmd.StringVar(&phone, "phone", "", "Phone number of the person")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteCmd.StringVar(&name, "name", "", "Name of the person")

	getCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	getCmd.StringVar(&name, "name", "", "Name of the person")

	if len(os.Args) < 2 {
		fmt.Println("Expected a sub-command -- add or delete or get.")
		usage()
		os.Exit(1)
	}

	// Parse command-line args
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if addCmd.NFlag() < 2 {
			fmt.Println("You must supply both --name and --phone options")
			addCmd.Usage()
			os.Exit(1)
		}
		newRecord := Person{
			Name:  name,
			Phone: phone,
		}
		records.AddEntry(newRecord)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if deleteCmd.NFlag() < 1 {
			fmt.Println("You must supply --name option")
			deleteCmd.Usage()
			os.Exit(1)
		}
		records.DeleteEntry(name)
	case "get":
		getCmd.Parse(os.Args[2:])
		entries := records.GetEntry(name)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
		fmt.Fprintln(w, "Name\tPhone\t")
		for _, v := range entries {
			fmt.Fprintln(w, v.Name, "\t", v.Phone, "\t")
		}
		w.Flush()
	default:
		fmt.Println("Invalid subcommand specified")
		usage()
		os.Exit(1)
	}

	// save the file
	file, _ := json.MarshalIndent(records.Entries, "", " ")
	err = ioutil.WriteFile(phoneBookPath, file, 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s", phoneBookPath)
		os.Exit(1)
	}
}
