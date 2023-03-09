package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	Title             string           `json:"title,omitempty" bson:"title" required:"true"`
	Description_short string           `json:"description_short,omitempty" bson:"description_short" required:"true"`
	Description_long  string           `json:"description_long,omitempty" bson:"description_long"`
	Date              time.Time        `json:"date,omitempty" bson:"date"`
	Organizer         string           `json:"organizer,omitempty" bson:"organizer" required:"true"`
	Place             string           `json:"place,omitempty" bson:"place"`
	State             string           `json:"state,omitempty" bson:"state"`
	SubscribersRef    []SubscribersRef `json:"-" bson:"subscribers,omitempty"`
}

type EventResponse struct {
	Id                primitive.ObjectID `json:"id" bson:"_id"`
	Title             string             `json:"title" bson:"title" required:"true"`
	Description_short string             `json:"description_short" bson:"description_short" required:"true"`
	Description_long  string             `json:"description_long" bson:"description_long"`
	Date              time.Time          `json:"date" bson:"date"`
	Organizer         string             `json:"organizer" bson:"organizer" required:"true"`
	Place             string             `json:"place" bson:"place"`
	State             string             `json:"state" bson:"state"`
	SubscribersRef    []SubscribersRef   `json:"-" bson:"subscribers,omitempty"`
}

type EventFilt struct {
	Title   string    `json:"title,omitempty" bson:"title,omitempty"`
	DateIn  time.Time `json:"dateIn,omitempty" bson:"dateIn,omitempty"`
	DateOut time.Time `json:"dateOut,omitempty" bson:"dateOut,omitempty"`
	State   string    `json:"state,omitempty" bson:"state,omitempty"`
}

type EventUpdate struct {
	Title             string           `json:"title,omitempty" bson:"title,omitempty"`
	Description_short string           `json:"description_short,omitempty" bson:"description_short,omitempty"`
	Description_long  string           `json:"description_long,omitempty" bson:"description_long,omitempty"`
	Date              time.Time        `json:"date,omitempty" bson:"date,omitempty"`
	Organizer         string           `json:"organizer,omitempty" bson:"organizer,omitempty"`
	Place             string           `json:"place,omitempty" bson:"place,omitempty"`
	State             string           `json:"state,omitempty" bson:"state,omitempty"`
	SubscribersRef    []SubscribersRef `json:"-" bson:"subscribers,omitempty"`
}
type SubscribersRef struct {
	ID   primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
}
