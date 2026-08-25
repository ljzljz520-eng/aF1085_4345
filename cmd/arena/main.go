package main

import (
	"example.com/arena/internal/api"
	"example.com/arena/internal/registry"
	"example.com/arena/internal/store"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	path := flag.String("db", "arena.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	db, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	reg := registry.New(db)
	mux := http.NewServeMux()
	mux.Handle("/records", api.New(reg))
	mux.HandleFunc("/health", api.Health)
	fmt.Println("arena listening", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
