package games

import (
	"net/http"

	"github.com/gorilla/mux"
	"tts_deck_build/internal/errors"
	"tts_deck_build/internal/games"
	"tts_deck_build/internal/utils"
)

// Request to update a game
//
// swagger:parameters RequestUpdateGame
type RequestUpdateGame struct {
	// In: path
	// Required: true
	Name string `json:"name"`
	// In: body
	// Required: true
	Body struct {
		games.UpdateGameRequest
	}
}

// Status of game update
//
// swagger:response ResponseUpdateGame
type ResponseUpdateGame struct {
}

// swagger:route PATCH /games/{name} Games RequestUpdateGame
//
// Update game
//
// Allows you to update an existing game
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: ResponseUpdateGame
//       default: ResponseError
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	req := &games.UpdateGameRequest{}
	err := utils.RequestToObject(r.Body, &req)
	if err != nil {
		errors.ResponseError(w, errors.InternalError.AddMessage(err.Error()))
		return
	}

	e := games.UpdateGame(name, req)
	if e != nil {
		errors.ResponseError(w, e)
	}
	return
}
