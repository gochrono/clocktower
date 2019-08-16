package api

import (
	"github.com/gochrono/castle/pkg/handlers"
	m "github.com/gochrono/castle/pkg/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq" // Load postges drivers
	"github.com/sirupsen/logrus"
	"os"
)

func Start(cfg *config.Configuration) error {
	var logger = logrus.New()
	logger.Out = os.Stdout

	// Open DB and defer close
	db, err := sqlx.Open("postgres", cfg.Database.DSN)
	if err != nil {
		return newSystemError("could not connect to database")
	}
	defer db.Close()

	router := httprouter.New()
	h := handlers.New(repository, logger)

	middlewareArgs := m.MiddlewareArguments{Repository: repository, Config: cfg, Logger: logger}

	// authenticate
	router.POST("/auth", m.Unauthenticated(h.AuthenticateUser, middlewareArgs))
	router.GET("/users/me", m.AuthWithToken(h.GetUserMe, middlewareArgs))
	router.GET("/users", m.AuthWithToken(h.GetUsers, middlewareArgs))
	router.POST("/users", m.AuthWithToken(h.CreateUser, middlewareArgs))
	router.DELETE("/users/:id", m.AuthWithToken(h.DeleteUser, middlewareArgs))
	router.PUT("/users/:id", m.AuthWithToken(h.UpdateUser, middlewareArgs))

	return nil
}
