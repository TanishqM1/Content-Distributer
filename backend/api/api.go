package api

import (
	"encoding/json"
	"net/http"
)

// this file is used as a helper file.
// it defines error messages that we will output in the project, as well as the structure for our json responses.

// struct of error responses.
// this struct will either be an internal error (program fault) or user error (something incorrect is passed in) -> though this error handling should be handled in the front too.
type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	// response header (error case)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// json for output (using Encode)
	json.NewEncoder(w).Encode(resp)
}

// wrapper function for writeError. This lets us use it with two different use cases.

var (
	// error is not internal.
	HandleRequestError = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// internal error. we log it speerately!
	HandleInternalError = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
