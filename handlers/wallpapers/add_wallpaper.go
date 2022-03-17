package wallpapers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Wallpaper struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}

func AddWallpaper(w http.ResponseWriter, r *http.Request) {
	var wallpaper Wallpaper
	_ = json.NewDecoder(r.Body).Decode(&wallpaper)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://username:pass@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("wallpapers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection.InsertOne(ctx, wallpaper)
	w.Write([]byte("message: Wallpaper added successfully"))

}
