package handler

import "net/http"

func (h Handler) addTravelHandler(w http.ResponseWriter, r *http.Request) error {
	return h.Service.Travel.AddTravel(r)
}