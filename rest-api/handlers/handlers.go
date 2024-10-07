package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sreejith1962/rest-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Collection *mongo.Collection
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := h.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	result, err := h.Collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		fmt.Println("Error")
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	result, err := h.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{Key: "$set", Value: book}},
	)

	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(result)

}
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	cursor, err := h.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)

	}
	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := h.Collection.InsertOne(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
