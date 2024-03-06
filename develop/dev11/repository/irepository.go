package repository

import (
	"dev11/models"
	"time"
)

type Repository interface {
	Create(models *models.Event) (*models.Event, error)
	Update(event *models.Event) (*models.Event, error)
	Delete(eventId uint) error
	GetByDay(day time.Time) ([]*models.Event, error)
	GetByWeek(firstDay time.Time, lastDay time.Time) ([]*models.Event, error)
	GetByMonth(month time.Time) ([]*models.Event, error)
}
