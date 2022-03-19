package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
	r.HandleFunc("/getwallpaper", wal.GetAllWallpaper).Methods("GET")
	r.HandleFunc("/delete/{id}", DeleteWallpaper).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func DeleteWallpaper(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tushardbanduser:Tushartxt11223344@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("wallpapers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"count":  strconv.Itoa(int(result.DeletedCount)),
	})
}
