package handlers

import (
	"encoding/json"
	"github.com/gochrono/castle/middleware"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	DetailInvalidUserID = "Invalid user ID"
)

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users, err := h.repository.GetUsers()
	if err != nil {
		h.logger.WithError(err).Warn("get users")
	}

	json.NewEncoder(w).Encode(users)
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIDRaw := ps.ByName("id")
	userID, err := strconv.Atoi(userIDRaw)

	if err != nil {
		h.logger.WithError(err).Warn("invalid user ID")
		h.SendError(w, http.StatusBadRequest, DetailInvalidUserID)
		return
	}

	h.logger.WithFields(log.Fields{
		"user": userID,
	}).Info("deleting user")

	user, err := h.repository.GetUserByID(userID)
	if err != nil {
		h.logger.WithError(err).Warn("fetch user")
		h.SendError(w, http.StatusBadRequest, DetailInternalServerError)
		return
	}

	err = h.repository.DeleteUserByID(userID)
	if err != nil {
		h.logger.WithError(err).Warn("delete user")
		h.SendError(w, http.StatusBadRequest, DetailInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIDRaw := ps.ByName("id")
	userID, err := strconv.Atoi(userIDRaw)

	if err != nil {
		h.logger.WithError(err).Warn("invalid user ID")
		h.SendError(w, http.StatusBadRequest, DetailInvalidUserID)
		return
	}

	h.logger.WithFields(log.Fields{
		"user": userID,
	}).Info("fetching user")

	user, err := h.repository.GetUserByID(userID)
	if err != nil {
		h.logger.WithError(err).Warn("fetch user")
		h.SendError(w, http.StatusBadRequest, DetailInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h Handler) GetUserMe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := middleware.GetCurrentUser(r.Context())
	json.NewEncoder(w).Encode(user)
}
