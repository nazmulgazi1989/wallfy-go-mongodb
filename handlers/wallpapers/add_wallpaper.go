package wallpapers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Wallpaper struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name"`

	Image     string `json:"image"`
	Desc      string `json:"desc"`
	Category  string `json:"category"`
	CreatedAt string `json:"createdAt"`
}

func AddWallpaper(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tushardbanduser:Tushartxt11223344@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("wallpapers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var wallpaper Wallpaper
	err = json.NewDecoder(r.Body).Decode(&wallpaper)
	if err != nil {
		log.Fatal(err)
	}

	wallpaper.CreatedAt = time.Now().String()
	objectId := primitive.NewObjectID()
	wallpaper.Id = objectId.Hex()
	result, err := collection.InsertOne(ctx, wallpaper)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)

}
