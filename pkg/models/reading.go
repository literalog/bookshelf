package models

import (
	"time"

	"github.com/google/uuid"
)

type Reading struct {
	Id          string    `json:"id" bson:"_id"`
	UserId      string    `json:"user_id" bson:"user_id"`
	BookId      string    `json:"book_id" bson:"book_id"`
	CurrentPage int       `json:"current_page" bson:"current_page"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

type ReadingRequest struct {
	UserId      string `json:"user_id" bson:"user_id"`
	BookId      string `json:"book_id" bson:"book_id"`
	CurrentPage int    `json:"current_page" bson:"current_page"`
}

func NewReading(req ReadingRequest) *Reading {
	return &Reading{
		Id:          uuid.NewString(),
		UserId:      req.UserId,
		BookId:      req.BookId,
		CurrentPage: req.CurrentPage,
		UpdatedAt:   time.Now(),
	}
}
