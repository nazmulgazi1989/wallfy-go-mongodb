package main

import (
	"log"
	"net/http"
	auth "tusharhow/wallpaper/handlers/auth"
	wal "tusharhow/wallpaper/handlers/wallpapers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/addwallpaper", wal.AddWallpaper).Methods("POST")
	r.HandleFunc("/getwallpaper", wal.GetAllWallpaper).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

