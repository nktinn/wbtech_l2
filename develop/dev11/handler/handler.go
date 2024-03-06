package handler

import (
	"dev11/calendar"
	"dev11/response"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type DefaultHandler struct {
	CalendarUC calendar.Calendar
}

func NewHandler(clndr calendar.Calendar) *DefaultHandler {
	return &DefaultHandler{CalendarUC: clndr}
}

func (h *DefaultHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendJsonResponse(w, 405, "only POST requests allowed")
		return
	}

	validEvent := response.CreateEventValidate(r)
	if validEvent == nil {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	eventWithID, err := h.CalendarUC.CreateEvent(validEvent)
	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, eventWithID)
}

func (h *DefaultHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendJsonResponse(w, 405, "only POST requests allowed")
		return
	}

	validEvent := response.UpdateEventValidate(r)
	if validEvent == nil {
		response.SendJsonResponse(w, 400, "ваши data invalid")
		return
	}

	updateEvent, err := h.CalendarUC.UpdateEvent(validEvent)
	errNotFound := fmt.Errorf("ID not found")
	if errors.Is(errNotFound, err) {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, updateEvent)
}

func (h *DefaultHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendJsonResponse(w, 405, "only POST requests allowed")
		return
	}

	eventId := response.DeleteEventValidate(r)
	if eventId == 0 {
		response.SendJsonResponse(w, 400, "ваши data invalid")
		return
	}

	err := h.CalendarUC.DeleteEvent(eventId)
	notFoundErr := fmt.Errorf("ID not found")
	if errors.Is(notFoundErr, err) {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, eventId)
}

func (h *DefaultHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.SendJsonResponse(w, 405, "only GET requests allowed")
		return
	}

	validDay, isValid := response.GetDayValidate(r)
	if !isValid {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	eventsForDay, err := h.CalendarUC.GetByDay(validDay)
	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, eventsForDay)
}

func (h *DefaultHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.SendJsonResponse(w, 405, "only GET requests allowed")
		return
	}

	firstDay, lastDay, isValid := response.GetWeekValidate(r)
	if !isValid {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	eventsForWeek, err := h.CalendarUC.GetByWeek(firstDay, lastDay)
	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, eventsForWeek)
}

func (h *DefaultHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.SendJsonResponse(w, 405, "only GET requests allowed")
		return
	}

	validMonth, isValid := response.GetMonthValidate(r)
	if !isValid {
		response.SendJsonResponse(w, 400, "data invalid")
		return
	}

	eventsForMonth, err := h.CalendarUC.GetByMonth(validMonth)
	if err != nil {
		response.SendJsonResponse(w, 503, "program error")
		return
	}

	response.SendJsonResponse(w, 200, eventsForMonth)
}

func (h *DefaultHandler) MiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf("%s %s %s Body: %s", r.Method, r.RemoteAddr, r.URL.Path, r.Form)
	}
}
