package handlers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"Appointy-Instagram/data"
	"Appointy-Instagram/functions"
)

type UserHandler struct {
	userCollection *mongo.Collection
}

func NewUserHandler(col *mongo.Collection) *UserHandler {
	return &UserHandler{
		userCollection: col,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			h.getUser(w, r)
		}
	case http.MethodPost:
		{
			h.createUser(w, r)
		}
	default:
		{
			http.Error(w, "Method not implemented", http.StatusMethodNotAllowed)
		}
	}
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &data.InUser{}
	ok := functions.ReadJson(w, r, user)
	if !ok {
		return
	}

	err1 := functions.ValidateUser(user)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x\n", hashedPassword.Sum(nil))

	userResult, err := h.userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Successfully created user with id: %v", userResult.InsertedID)))
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	fmt.Println(id)

	user := &data.OutUser{}
	userResult := h.userCollection.FindOne(context.Background(), bson.D{{"_id", id}})
	err := userResult.Decode(user)
	if err != nil {
		w.Write([]byte("unable to get data"))
	} else {
		functions.WriteJson(w, r, user)
	}

}
