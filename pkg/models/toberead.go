package models

import (
	"time"

	"github.com/google/uuid"
)

type ToBeRead struct {
	Id        string    `json:"id" bson:"_id"`
	UserId    string    `json:"user_id" bson:"user_id"`
	BookId    string    `json:"book_id" bson:"book_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type ToBeReadRequest struct {
	UserId string `json:"user_id" bson:"user_id"`
	BookId string `json:"book_id" bson:"book_id"`
}

func NewToBeRead(req ToBeReadRequest) *ToBeRead {
	return &ToBeRead{
		Id:        uuid.NewString(),
		UserId:    req.UserId,
		BookId:    req.BookId,
		CreatedAt: time.Now(),
	}
}
