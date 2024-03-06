package calendar

import (
	"dev11/models"
	"time"
)

type Calendar interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	UpdateEvent(event *models.Event) (*models.Event, error)
	DeleteEvent(eventId uint) error
	GetByDay(day time.Time) ([]*models.Event, error)
	GetByWeek(firstDay time.Time, lastDay time.Time) ([]*models.Event, error)
	GetByMonth(month time.Time) ([]*models.Event, error)
}
