package calendar

import (
	"dev11/models"
	"dev11/repository"
	"errors"
	"fmt"
	"time"
)

type CalendarRepo struct {
	repository repository.Repository
}

func NewCalendar(repository repository.Repository) *CalendarRepo {
	return &CalendarRepo{repository: repository}
}

func (c *CalendarRepo) CreateEvent(event *models.Event) (*models.Event, error) {
	createdEvent, err := c.repository.Create(event)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return createdEvent, nil
}

func (c *CalendarRepo) UpdateEvent(event *models.Event) (*models.Event, error) {
	updatedEvent, err := c.repository.Update(event)
	errNotFound := fmt.Errorf("ID not found")
	if errors.Is(errNotFound, err) {
		return nil, errNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return updatedEvent, nil
}

func (c *CalendarRepo) DeleteEvent(eventId uint) error {
	err := c.repository.Delete(eventId)
	errNotFound := fmt.Errorf("ID not found")
	if errors.Is(errNotFound, err) {
		return errNotFound
	}
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}

	return nil
}

func (c *CalendarRepo) GetByDay(day time.Time) ([]*models.Event, error) {
	forDay, err := c.repository.GetByDay(day)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forDay, nil
}

func (c *CalendarRepo) GetByWeek(firstDay time.Time, lastDay time.Time) ([]*models.Event, error) {
	forWeek, err := c.repository.GetByWeek(firstDay, lastDay)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forWeek, nil
}

func (c *CalendarRepo) GetByMonth(month time.Time) ([]*models.Event, error) {
	forMonth, err := c.repository.GetByMonth(month)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}

	return forMonth, nil
}
