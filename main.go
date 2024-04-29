package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/vitoschuster/events/handlers"
	"github.com/vitoschuster/events/store"
)

func main() {
	cfg := mysql.Config{
		User:                 store.Envs.DBUser,
		Passwd:               store.Envs.DBPassword,
		Addr:                 store.Envs.DBAddress,
		DBName:               store.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := store.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStore(db)

	initStorage(db)

	router := mux.NewRouter()

	handler := handlers.New(store)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")
	router.HandleFunc("/events", handler.HandleListEvents).Methods("GET")
	router.HandleFunc("/events", handler.HandleAddEvent).Methods("POST")
	router.HandleFunc("/events/{id}", handler.HandleDeleteEvent).Methods("DELETE")
	router.HandleFunc("/events/search", handler.HandleSearchEvent).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	http.ListenAndServe(":8080", router)
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Database!")
}
