package collections

import (
	"net/http"

	"github.com/gorilla/mux"

	"tts_deck_build/internal/collections"
	"tts_deck_build/internal/network"
)

// Requesting a list of existing collections
//
// swagger:parameters RequestListOfCollections
type RequestListOfCollections struct {
	// In: path
	// Required: true
	Game string `json:"game"`
	// In: query
	// Required: false
	Sort string `json:"sort"`
}

// List of collections
//
// swagger:response ResponseListOfCollections
type ResponseListOfCollections struct {
	// In: body
	// Required: true
	Body struct {
		// Required: true
		Data []*collections.CollectionInfo `json:"data"`
	}
}

// swagger:route GET /games/{game}/collections Collections RequestListOfCollections
//
// Get collections list
//
// Get a list of existing collections
// Sort values: name, name_desc, created, created_desc
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
//       200: ResponseListOfCollections
//       default: ResponseError
func ListHandler(w http.ResponseWriter, r *http.Request) {
	gameID := mux.Vars(r)["game"]
	sort := r.URL.Query().Get("sort")
	items, e := collections.NewService().List(gameID, sort)
	if e != nil {
		network.ResponseError(w, e)
		return
	}
	network.Response(w, items)
}
