package main

import (
	"embed"
	"github.com/sirupsen/logrus"
	"html/template"
)

//go:embed views/*
var htmlFiles embed.FS

//go:embed static/*
var staticFiles embed.FS

func main() {

	config := NewAppConfig()
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	templates, err := template.ParseFS(htmlFiles, "views/*.html")
	if err != nil {
		logger.WithError(err).Panic("Error Loading Templates!")
	}

	s := NewServer(config, logger, templates)

	s.Initilize()
	s.Listen()

}
