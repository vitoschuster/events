package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitoschuster/events/components"
	"github.com/vitoschuster/events/types"
	"github.com/vitoschuster/events/views"
)

func (h *Handler) HandleListEvents(w http.ResponseWriter, r *http.Request) {
	isAddingEvent := r.URL.Query().Get("isAddingEvent") == "true"

	events, err := h.store.GetEvents()
	if err != nil {
		log.Println(err)
		return
	}

	views.Events(events, isAddingEvent).Render(r.Context(), w)
}

func (h *Handler) HandleAddEvent(w http.ResponseWriter, r *http.Request) {
	event := &types.Event{
		Location:    r.FormValue("name"),
		EventTime:   r.FormValue("event_time"),
		Name:        r.FormValue("location"),
		Description: r.FormValue("description"),
		Price:       r.FormValue("price"),
		ImageURL:    r.FormValue("imageURL"),
	}

	newEvent, err := h.store.CreateEvent(event)
	if err != nil {
		log.Println(err)
		return
	}

	components.EventTile(newEvent).Render(r.Context(), w)
}

func (h *Handler) HandleSearchEvent(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("search")

	events, err := h.store.FindEventsByLocationNameOrDesc(text)
	if err != nil {
		log.Println(err)
		return
	}

	components.EventsList(events).Render(r.Context(), w)
}

func (h *Handler) HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.store.DeleteEvent(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
