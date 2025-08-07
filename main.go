package main

import "log"

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatalln(err)
	}

	if err := store.Init(); err != nil {
		log.Fatalln(err)
	}

	server := NewApiServer(":6969", store)
	server.Run()
}
