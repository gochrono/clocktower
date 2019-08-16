package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gochrono/clocktower/config"
	"github.com/gochrono/clocktower/models"
	"github.com/julienschmidt/httprouter"
)

var (
	DetailInvalidAuthorizationHeader = "Invalid or missing Authorization header"
	DetailUserNotFound               = "User not found"
	DetailTokenExpired               = "Token is expired"
	DetailMalformedToken             = "Malformed JWT token (claims)"
	DetailUserCreationFailed         = "User creation failed"
	DetailUserSelectionFailed        = "User selection failed"
	DetailUserProfileRetrievalFailed = "User profile retrieval failed"
)

type MiddlewareArguments struct {
	Repository models.Repository
	Config     config.Config
	Logger     *logrus.Logger
}

func AuthWithToken(h httprouter.Handle, args MiddlewareArguments) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			args.Logger.WithFields(logrus.Fields{
				"authorization_header": authHeader,
			}).Warn("bad authorization header")
			SendError(w, http.StatusBadRequest, DetailInvalidAuthorizationHeader)
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.SecretKey(), nil
		})

		if err != nil {
			args.Logger.WithError(err).Warn("validating claim")
			if err == jwt.ErrSignatureInvalid {
				SendError(w, http.StatusUnauthorized, "invalid signature")
			} else {
				SendError(w, http.StatusBadRequest, "bad request")
			}
			return
		}

		if !token.Valid {
			args.Logger.Warn("invalid token")
			SendError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), ContextCurrentUser, &models.User{})
		h(w, r.WithContext(ctx), ps)
	}
}
