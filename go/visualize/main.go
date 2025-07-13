package main

import (
	"encoding/json"
	"fmt"
	goforj_godump "github.com/goforj/godump"
	"github.com/yassinebenaid/godump"
	"gopkg.in/yaml.v2"
)

type Author struct {
	Name string
	Age  int
}

type Genre struct {
	Name        string
	Description string
}

type Book struct {
	internalID int
	Title      string
	Author     *Author
	Genre      *Genre
	Price      float64
	ISBN       string
	Quantity   int
}

type Bookstore struct {
	Name    string
	Address string
	Books   []*Book
}

func NewSampleBookstore() *Bookstore {
	programmingGenre := &Genre{
		Name:        "Programming",
		Description: "Books related to programming and software development.",
	}

	return &Bookstore{
		Name:    "Scott's Computer Book Emporium",
		Address: "123 Fictional Street",
		Books: []*Book{
			{
				internalID: 1,
				Title:      "The Go Programming Language",
				Author: &Author{
					Name: "Alan A. A. Donovan",
					Age:  45,
				},
				Genre:    programmingGenre,
				Price:    39.99,
				ISBN:     "9780134190440",
				Quantity: 10,
			},
			{
				internalID: 2,
				Title:      "Clean Code",
				Author: &Author{
					Name: "Robert C. Martin",
					Age:  65,
				},
				Genre:    programmingGenre,
				Price:    29.99,
				ISBN:     "9780132350884",
				Quantity: 5,
			},
			{
				internalID: 3,
				Title:      "The Pragmatic Programmer",
				Author: &Author{
					Name: "Andrew Hunt",
					Age:  50,
				},
				Genre:    programmingGenre,
				Price:    34.99,
				ISBN:     "9780201616224",
				Quantity: 8,
			},
		},
	}
}

func vis1(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 1 (print using \"%%v\"): Bookstore\n")
	fmt.Printf("%v\n", bookstore)
	fmt.Printf("Visualization 1 (print using \"%%v\"): A single book\n")
	fmt.Printf("%v\n", bookstore.Books[0])
}

func vis2(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 2 (print using \"%%+v\"): Bookstore\n")
	fmt.Printf("%+v\n", bookstore)
	fmt.Printf("Visualization 2 (print using \"%%+v\"): A single book\n")
	fmt.Printf("%+v\n", bookstore.Books[0])
}

func vis3(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 4 (print using json):\n")

	data, err := json.Marshal(bookstore)
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

func vis4(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 4 (print using json, indented):\n")

	data, err := json.MarshalIndent(bookstore, "", "  ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

func vis5(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 5 (print using yaml):\n")

	data, err := yaml.Marshal(bookstore)
	if err != nil {
		fmt.Printf("Error converting to YAML: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

func vis6(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 6 (print using github.com/yassinebenaid/godump):\n")

	godump.Dump(bookstore)
}

func vis7(bookstore *Bookstore) {
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Visualization 7 (print using github.com/goforj/godump):\n")

	goforj_godump.Dump(bookstore)
}

func main() {
	bookstore := NewSampleBookstore()
	vis1(bookstore)
	vis2(bookstore)
	vis3(bookstore)
	vis4(bookstore)
	vis5(bookstore)
	vis6(bookstore)
	vis7(bookstore)
}
