package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/URAmiRBin/go-learning/login/config/db"
	"github.com/URAmiRBin/go-learning/login/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)

	var result model.ResponseResult
	if err != nil {
		fmt.Println("ERROR Unmarshaling")
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	collection, err := db.GetDBCollection()

	if err != nil {
		fmt.Println("ERROR get db collection")
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	var res model.User
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&res)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			fmt.Println("No user found, registering")
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
			if err != nil {
				result.Error = "Error while hashing, try again"
				json.NewEncoder(w).Encode(result)
				return
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				result.Error = "Error while creating user"
				json.NewEncoder(w).Encode(result)
				return
			}

			result.Result = "Register successful"
			json.NewEncoder(w).Encode(result)
			return
		}
		fmt.Println("DB Error but not no document found")
		result.Error = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}
	result.Result = "Username already exists"
	json.NewEncoder(w).Encode(result)
	return
}
