package main

import ()

func Main() {
	cfg := config.ReadConfig("castle.toml")
	logger.WithFields(logrus.Fields{
		"database": cfg.Database,
	}).Info("loaded cfg")
}
