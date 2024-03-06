package handler

import "net/http"

type Handler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)

	GetEventsForDay(w http.ResponseWriter, r *http.Request)
	GetEventsForWeek(w http.ResponseWriter, r *http.Request)
	GetEventsForMonth(w http.ResponseWriter, r *http.Request)

	MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc
}
