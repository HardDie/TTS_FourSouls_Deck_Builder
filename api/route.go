package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"tts_deck_build/api/cards"
	"tts_deck_build/api/collections"
	"tts_deck_build/api/decks"
	"tts_deck_build/api/games"
	"tts_deck_build/api/images"
	"tts_deck_build/api/settings"
	"tts_deck_build/api/web"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)
	web.Init(routes)
	games.Init(routes)
	collections.Init(routes)
	decks.Init(routes)
	cards.Init(routes)
	images.Init(routes)
	settings.Init(routes)
	routes.Use(corsMiddleware)
	return routes
}

// CORS headers
func corsSetupHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// CORS Headers middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corsSetupHeaders(w)
		next.ServeHTTP(w, r)
	})
}
