package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, password string) *Account {
	acc, err := NewAccount(fname, lname, password)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New account number :=> ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Basir", "SecondAccount", "hunter888")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)

	server.Run()
	//
	fmt.Print("The server is run sucessfully")
}
