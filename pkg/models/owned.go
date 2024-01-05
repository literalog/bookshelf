package models

import "github.com/google/uuid"

type Owned struct {
	Id     string `json:"id" bson:"_id"`
	UserId string `json:"user_id" bson:"user_id"`
	BookId string `json:"book_id" bson:"book_id"`
}

type OwnedRequest struct {
	UserId string `json:"user_id" bson:"user_id"`
	BookId string `json:"book_id" bson:"book_id"`
}

func NewOwned(req OwnedRequest) Owned {
	return Owned{
		Id:     uuid.NewString(),
		UserId: req.UserId,
		BookId: req.BookId,
	}
}
