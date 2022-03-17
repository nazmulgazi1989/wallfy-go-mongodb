package login


import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tushardbanduser:Tushartxt11223344@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
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
		if users[0].Password == user.Password {

			token, err := generateToken(users[0].ID)
			if err != nil {
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(map[string]string{
				"status": "success",
				"token":  token})
		} else {
			w.Write([]byte("message: Incorrect password"))
		}
	} else {
		w.Write([]byte("message: User does not exist"))
	}

}
func generateToken(userId string) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "tushsshss") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	t, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}