package usecase

import (
	"errors"
	"time"

	"github.com/brunoseba/event-api/internal/entity"
	"github.com/brunoseba/event-api/internal/repository"
	"github.com/brunoseba/event-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventUseCase struct {
	eventRepo repository.EventRepositoryImp
	//eventRepo repository.EventRepository
}

func NewEventUseCase(eventRepo repository.EventRepositoryImp) EventUseCase {
	return EventUseCase{eventRepo}
}

func (eu *EventUseCase) CreateEventService(event *entity.Event) (*entity.EventResponse, error) {
	//dateToTime, _ := time.Parse(time.RFC3339, event.Date) //time.RFC3339
	newEvent := &entity.Event{
		Title:             event.Title,
		Description_short: event.Description_short,
		Description_long:  event.Description_long,
		Date:              event.Date,
		Organizer:         event.Organizer,
		Place:             event.Place,
		State:             event.State,
	}

	eventRes, err := eu.eventRepo.CreateEvent(newEvent)
	if err != nil {
		return nil, err
	}
	return eventRes, nil
}

func (eu *EventUseCase) GetEvent(id string) (*entity.EventResponse, error) {
	event, err := eu.eventRepo.GetEventByID(id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (eu *EventUseCase) GetAllEventService() ([]primitive.M, error) {
	event, err := eu.eventRepo.GetAllEvent()
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (eu *EventUseCase) GetEventFiltDateService(isAdmin bool, dIn, dOut time.Time, title, state string) ([]*entity.EventResponse, error) {
	valid := utils.ValidateState(isAdmin, state)
	isDate := utils.ValidateDate(dIn, dOut)

	if (valid != utils.Draft) && isDate {
		events, err := eu.eventRepo.GetEventByDate(dIn, dOut, valid)
		if err != nil {
			return nil, err
		}
		return events, nil
	}
	events, err := eu.eventRepo.GetEventState(valid)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (eu *EventUseCase) UpdateEvent(id string, event *entity.EventUpdate) (*entity.EventResponse, error) {

	eventUpdated, err := eu.eventRepo.UpdateEvent(id, event)
	if err != nil {
		return nil, err
	}

	return eventUpdated, nil
}

func (eu *EventUseCase) AddEventToUser(userId string, eventID string) error {

	event, er := eu.eventRepo.GetEventByID(eventID)
	if er != nil {
		return er
	}

	date := event.Date
	today := time.Now()
	if date.After(today) {
		err := eu.eventRepo.SubscriteToEvent(userId, eventID)
		if err != nil {
			return err
		}
	}
	return errors.New("can't suscribe to Event, the event ended")
}

func (eu *EventUseCase) GetListEventUser(userId string) ([]*entity.EventResponse, error) {
	eventos, err := eu.eventRepo.GetEventsByUser(userId)
	if err != nil {
		return nil, err
	}
	return eventos, nil
}

func (eu *EventUseCase) DeleteEventById(id string) error {
	err := eu.eventRepo.DeleteEvent(id)
	if err != nil {
		return err
	}
	return nil
}
