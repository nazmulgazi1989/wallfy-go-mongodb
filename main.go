package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	auth "tusharhow/wallpaper/handlers/auth"
	wal "tusharhow/wallpaper/handlers/wallpapers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/addwallpaper", wal.AddWallpaper).Methods("POST")
	r.HandleFunc("/getwallpaper", GetAllWallpaper).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type AllWallpaper struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
	Uploader string `json:"uploader"`


}

func GetAllWallpaper(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tushardbanduser:Tushartxt11223344@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("wallpapers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var allWallpapers []AllWallpaper
	err = result.All(ctx, &allWallpapers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allWallpapers)
	json.NewEncoder(w).Encode(allWallpapers)
}