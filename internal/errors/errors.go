package errors

import (
	"net/http"

	"tts_deck_build/internal/utils"
)

const (
	// FileNotExist =
	// FileExists   = "file_exists"
	UnknownError = "unknown_error"
	Done         = "done"
)

var (
	InternalError = NewError("internal error").HTTP(http.StatusInternalServerError)
	DataInvalid   = NewError("game data invalid").HTTP(http.StatusNoContent)
	GameNotExists = NewError("game not exists")
	FileExists    = NewError("file exists")
	// FileNotExist = NewError("file not exist")
)

type Error struct {
	Message string `json:"message"`
	code    int
}

func NewError(message string) *Error {
	return &Error{
		Message: message,
		code:    http.StatusBadRequest,
	}
}

func (e *Error) HTTP(code int) *Error {
	e.code = code
	return e
}

func (e Error) AddMessage(message string) *Error {
	e.Message += ": " + message
	return &e
}

func ResponseError(w http.ResponseWriter, e *Error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.code)
	if len(e.Message) > 0 {
		_, err := w.Write(utils.ToJson(e))
		utils.IfErrorLog(err)
	}
	return
}
