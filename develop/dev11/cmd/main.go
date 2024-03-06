package main

import (
	calendar "dev11/calendar"
	"dev11/config"
	"dev11/handler"
	"dev11/repository"
	"net/http"
)

func main() {
	cfg := config.GetConfig()

	rep := repository.NewMapRepository()

	CalendarRepo := calendar.NewCalendar(rep)

	h := handler.NewHandler(CalendarRepo)

	http.HandleFunc("/create_event", h.MiddlewareLogger(h.CreateEvent))
	http.HandleFunc("/update_event", h.MiddlewareLogger(h.UpdateEvent))
	http.HandleFunc("/delete_event", h.MiddlewareLogger(h.DeleteEvent))
	http.HandleFunc("/events_for_day", h.MiddlewareLogger(h.GetEventsForDay))
	http.HandleFunc("/events_for_week", h.MiddlewareLogger(h.GetEventsForWeek))
	http.HandleFunc("/events_for_month", h.MiddlewareLogger(h.GetEventsForMonth))

	err := http.ListenAndServe(cfg.HttpServerPort, nil)
	if err != nil {
		return
	}
}
