package main

import ()

func Main() {
	cfg := config.ReadConfig("clocktower.toml")
	logger.WithFields(logrus.Fields{
		"database": cfg.Database,
	}).Info("loaded cfg")
}
