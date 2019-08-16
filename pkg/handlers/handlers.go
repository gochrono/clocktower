// Package handlers contains the application logic (controllers).
//
// All handlers expect JSON content (if any) and return JSON content, along
// with proper status codes. More information at:
// http://docs.crickapi.apiary.io/.
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gochrono/clocktower/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var (
	// DefaultContentType is the default content type for handler responses.
	DefaultContentType = "application/json"

	// DetailInvalidRequest is the error message used when the request parsing has failed.
	DetailInvalidRequest = "Invalid request"
	// DetailMalformedJSON is the error message used when it is not possible to parse the JSON content.
	DetailMalformedJSON = "Malformed JSON"

	// DetailUserIsNotAllowedToPerformOperation is the error message when the user cannot perform an operation.
	DetailUserIsNotAllowedToPerformOperation = "Not allowed"

	// DetailInvalidCredentials is the error message when the given login credentials are wrong
	DetailInvalidCredentials = "Invalid Credentials"

	DetailInternalServerError = "Internal server error"
)

// Handler is the structure that contains the different HTTP handlers.
type Handler struct {
	repository models.Repository
	logger     *logrus.Logger
}

// New creates the main handler.
func New(repository models.Repository, logger *logrus.Logger) Handler {
	return Handler{
		repository: repository,
		logger:     logger,
	}
}

func (h Handler) BindJSON(w http.ResponseWriter, r *http.Request, t interface{}) error {
	w.Header().Set("Content-Type", DefaultContentType)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.WithError(err).Warn("reading request body")
		h.SendError(w, http.StatusBadRequest, DetailInvalidRequest)
		return err
	}

	if err := json.Unmarshal(body, t); err != nil {
		h.logger.WithError(err).Warn("unmarshal body into json")
		h.SendError(w, http.StatusBadRequest, DetailMalformedJSON)
		return err
	}
	return nil
}

// SendError returns a HTTP error in JSON.
func (h Handler) SendError(w http.ResponseWriter, statusCode int, detail string) {
	w.Header().Set("Content-Type", DefaultContentType)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"title":  http.StatusText(statusCode),
		"detail": detail,
	})
}

func getIntOrDefault(value string, defaultValue int) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		v = defaultValue
	}

	return v
}

func getStringSlice(value string) []string {
	var slice []string

	if value == "" {
		return slice
	}

	for _, v := range strings.Split(value, ",") {
		slice = append(slice, strings.TrimSpace(v))
	}

	return slice
}
