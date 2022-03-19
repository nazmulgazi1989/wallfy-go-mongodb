package wallpapers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AllWallpaper struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	Desc     string `json:"desc"`
	Uploader string `json:"uploader"`
	Category string `json:"category"`
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
