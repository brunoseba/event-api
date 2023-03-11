package repository

import (
	"context"
	"errors"
	"time"

	"github.com/brunoseba/event-api/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventRepositoryImp struct {
	eventCollection *mongo.Collection
	ctx             context.Context
}

func NewEventReposiroty(eventCollection *mongo.Collection, ctx context.Context) EventRepositoryImp {
	return EventRepositoryImp{eventCollection, ctx}
}

func (er *EventRepositoryImp) CreateEvent(event *entity.Event) (*entity.EventResponse, error) {
	resp, err := er.eventCollection.InsertOne(er.ctx, event)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("the user is already exist")
		}
		return nil, err
	}

	var newEvent *entity.EventResponse
	query := bson.M{"_id": resp.InsertedID}

	if err = er.eventCollection.FindOne(er.ctx, query).Decode(&newEvent); err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (er *EventRepositoryImp) GetEventByID(id string) (*entity.EventResponse, error) {
	eventId, _ := primitive.ObjectIDFromHex(id)
	resp := &entity.EventResponse{}
	query := bson.M{"_id": eventId}

	err := er.eventCollection.FindOne(er.ctx, query).Decode(&resp)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &entity.EventResponse{}, err
		}
		return nil, err
	}
	return resp, nil
}

func (er *EventRepositoryImp) GetAllEvent() ([]primitive.M, error) {

	allEvent, err := er.eventCollection.Find(er.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer allEvent.Close(er.ctx)

	var eventos []bson.M
	if err = allEvent.All(er.ctx, &eventos); err != nil {
		return nil, err
	}
	return eventos, nil
}

func (er *EventRepositoryImp) GetEventByDate(dateStart, dateEnd time.Time, state string) ([]*entity.EventResponse, error) {
	var events []*entity.EventResponse
	filter := bson.D{
		{Key: "state", Value: "publicado"},
		{Key: "$and", Value: []bson.M{
			{"date": bson.D{{Key: "$gte", Value: dateStart}, {Key: "$lte", Value: dateEnd}}},
		}},
	}

	cursor, err := er.eventCollection.Find(er.ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(er.ctx)

	for cursor.Next(er.ctx) {
		event := &entity.EventResponse{}
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (er *EventRepositoryImp) GetEventState(state string) ([]*entity.EventResponse, error) {
	var events []*entity.EventResponse
	filter := bson.D{
		{Key: "state", Value: state},
	}

	cursor, err := er.eventCollection.Find(er.ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(er.ctx)

	for cursor.Next(er.ctx) {
		event := &entity.EventResponse{}
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (er *EventRepositoryImp) UpdateEvent(id string, event *entity.EventUpdate) (*entity.EventResponse, error) {

	eventId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: eventId}}
	update := bson.D{{Key: "$set", Value: event}}

	resp := er.eventCollection.FindOneAndUpdate(er.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After))

	var updateEvent *entity.EventResponse
	if err := resp.Decode(&updateEvent); err != nil {
		return nil, errors.New("could not update event")
	}
	return updateEvent, nil
}

func (er *EventRepositoryImp) GetEventsByUser(userId string) ([]*entity.EventResponse, error) {
	var events []*entity.EventResponse

	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	cursor, err := er.eventCollection.Find(er.ctx, bson.M{"SubscribersRef": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(er.ctx)

	for cursor.Next(er.ctx) {
		event := &entity.EventResponse{}
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (er *EventRepositoryImp) SubscriteToEvent(userID string, eventID string) error {
	userOID, _ := primitive.ObjectIDFromHex(userID)
	eventOID, _ := primitive.ObjectIDFromHex(eventID)

	filter := bson.M{"_id": eventOID}
	update := bson.M{"$addToSet": bson.M{"SubscribersRef": userOID}}

	er.eventCollection.UpdateOne(er.ctx, filter, update)

	return nil
}

func (er *EventRepositoryImp) DeleteEvent(eventID string) error {
	objectID, _ := primitive.ObjectIDFromHex(eventID)

	filter := bson.M{"_id": objectID}

	_, err := er.eventCollection.DeleteOne(er.ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
