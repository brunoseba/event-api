package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Nickname string `json:"nickname" bson:"nickname"`
	IsAdmin  bool   `bson:"isAdmin" default:"false"`
}

type UserResponse struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Nickname string             `json:"nickname" bson:"nickname"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin"`
}

type UserUpdate struct {
	Nickname string `json:"nickname,omitempty" bson:"nickname,omitempty"`
	IsAdmin  bool   `json:"isAdmin,omitempty" bson:"isAdmin,omitempty"`
}

type UserEvent struct {
	UserId  string `json:"id"`
	EventId string `json:"eventId"`
}
