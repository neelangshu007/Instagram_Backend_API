package handlers

import (
	"Appointy-Instagram/data"
	"Appointy-Instagram/functions"
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostUserHandler struct {
	postCollection *mongo.Collection
}

func NewPostUserHandler(col *mongo.Collection) *PostUserHandler {
	return &PostUserHandler{
		postCollection: col,
	}
}

func (h *PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{

			userId := r.URL.Path[len("/posts/users/"):]

			limit, offset, err1 := functions.GetLimitAndOffset(w, r)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
			}

			postCursor, err := h.postCollection.Find(context.Background(), bson.D{{"userId", userId}}, &options.FindOptions{
				Limit: &limit,
				Skip:  &offset,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			posts := &[]data.OutPost{}
			err = postCursor.All(context.Background(), posts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			functions.WriteJson(w, r, posts)
		}

	default:
		{
			w.Write([]byte("Method not implemented"))
		}
	}

}
