package repository

import (
	"dev11/models"
	"fmt"
	"sync"
	"time"
)

type MapRepository struct {
	mapEvents map[uint]*models.Event
	lastId    uint
	sm        *sync.Mutex
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		mapEvents: map[uint]*models.Event{},
		lastId:    0,
		sm:        &sync.Mutex{},
	}
}

func (m *MapRepository) Create(event *models.Event) (*models.Event, error) {
	m.sm.Lock()

	m.lastId++
	event.ID = m.lastId
	m.mapEvents[m.lastId] = event

	m.sm.Unlock()

	return event, nil
}

func (m *MapRepository) Update(event *models.Event) (*models.Event, error) {
	m.sm.Lock()

	e, ok := m.mapEvents[event.ID]
	if !ok {
		err := fmt.Errorf("ID not found")
		return nil, err
	}

	e.Name = event.Name
	e.Description = event.Description
	e.Date = event.Date

	m.mapEvents[event.ID] = e

	m.sm.Unlock()

	return e, nil
}

func (m *MapRepository) Delete(eventId uint) error {
	m.sm.Lock()
	_, ok := m.mapEvents[eventId]
	if !ok {
		err := fmt.Errorf("ID not found")
		return err
	}

	delete(m.mapEvents, eventId)
	m.sm.Unlock()

	return nil
}

func (m *MapRepository) GetByDay(day time.Time) ([]*models.Event, error) {
	m.sm.Lock()
	eventsByDay := make([]*models.Event, 0, 2)

	for _, e := range m.mapEvents {
		if e.Date.Year() == day.Year() && e.Date.Month() == day.Month() && e.Date.Day() == day.Day() {
			eventsByDay = append(eventsByDay, e)
		}
	}
	m.sm.Unlock()
	return eventsByDay, nil
}

func (m *MapRepository) GetByWeek(firstDay time.Time, lastDay time.Time) ([]*models.Event, error) {
	m.sm.Lock()
	eventsByWeek := make([]*models.Event, 0, 2)

	for _, e := range m.mapEvents {
		if e.Date.Unix() >= firstDay.Unix() && e.Date.Unix() < lastDay.Unix() {
			eventsByWeek = append(eventsByWeek, e)
		}
	}

	m.sm.Unlock()

	return eventsByWeek, nil
}

func (m *MapRepository) GetByMonth(month time.Time) ([]*models.Event, error) {
	m.sm.Lock()
	eventsByMonth := make([]*models.Event, 0, 2)

	for _, e := range m.mapEvents {
		if e.Date.Year() == month.Year() && e.Date.Month() == month.Month() {
			eventsByMonth = append(eventsByMonth, e)
		}
	}

	m.sm.Unlock()

	return eventsByMonth, nil
}
