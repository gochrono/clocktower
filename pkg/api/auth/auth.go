package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gochrono/clocktower/config"
	"github.com/gochrono/clocktower/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func (h Handler) AuthenticateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.logger.Info("fetching users")
	w.Header().Set("Content-Type", DefaultContentType)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.WithError(err).Warn("bulk insert frames")
		h.SendError(w, http.StatusBadRequest, DetailInvalidRequest)
		return
	}

	var login models.Login
	if err := json.Unmarshal(body, &login); err != nil {
		h.logger.WithError(err).Warn("unmarshal JSON frames to synchronize")
		h.SendError(w, http.StatusBadRequest, DetailMalformedJSON)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"username": login.Username,
		"password": login.Password,
	}).Info("login object")

	user, err := h.repository.GetUserByUsername(login.Username)
	if err != nil {
		h.logger.WithError(err).Warn("finding user")
	}

	if user.Password == login.Password {
		expirationTime := time.Now().Add(6 * time.Hour)
		claims := &models.Claims{
			Username: user.Username,
			Roles:    []string{"users:create"},
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(config.SecretKey())
		if err != nil {
			h.logger.WithError(err).Warn("signing token")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": tokenString,
		})
	} else {
		h.SendError(w, http.StatusUnauthorized, DetailInvalidCredentials)
	}
}
