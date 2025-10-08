package util

import (
	"log"
	"net/http"
)

// HandleError is a utility function to handle errors and respond with appropriate HTTP status codes.
func HandleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	http.Error(w, err.Error(), statusCode)
}