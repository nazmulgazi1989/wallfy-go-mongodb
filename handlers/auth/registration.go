package login

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

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://username:pass@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	result, err := collection.Find(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = result.All(ctx, &users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
	if len(users) > 0 {
		w.Write([]byte("message: User already exists"))
	} else {
		collection.InsertOne(ctx, user)

		w.Write([]byte("message: User registered successfully"))
	}

}
