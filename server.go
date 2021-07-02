package main

import (
	"WEB/home"
	"WEB/student"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type Server struct {
	config    *AppConfig
	PORT      int
	ADDRESS   string
	logger    *logrus.Logger
	Router    *mux.Router
	DB        *gorm.DB
	templates *template.Template
}

func NewServer(config *AppConfig, logger *logrus.Logger, templates *template.Template) *Server {
	return &Server{
		config:    config,
		PORT:      config.PORT,
		ADDRESS:   config.ADDRESS,
		logger:    logger,
		Router:    mux.NewRouter(),
		templates: templates,
	}
}

func (s *Server) addrouter() {

	s.Router.StrictSlash(true)

	sr := student.NewRepo(s.DB)
	sh := student.NewHandler(s.logger, sr)
	hh := home.NewHandler(s.logger, s.templates)

	s.Router.HandleFunc("/student", sh.Create).Methods(http.MethodPost)
	s.Router.HandleFunc("/student", sh.GetAll).Methods(http.MethodGet)
	s.Router.HandleFunc("/student/{roll}", sh.Update).Methods(http.MethodPost)
	s.Router.HandleFunc("/student/{roll}", sh.Delete).Methods(http.MethodDelete)
	s.Router.HandleFunc("/student/{roll}", sh.Get).Methods(http.MethodGet)

	s.Router.PathPrefix("/static").Handler(http.FileServer(http.FS(staticFiles)))
	s.Router.HandleFunc("/", hh.Home).Methods(http.MethodGet)
	s.Router.HandleFunc("/submit", hh.Submit).Methods(http.MethodPost)
}

func (s *Server) Initilize() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		s.config.DBHost,
		s.config.DBUser,
		s.config.DBPass,
		s.config.DBName,
		s.config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		s.logger.WithError(err).Error("Failed Connecting to DATABASE!")
	}

	s.logger.Info("Connected to DATABASE!")

	s.DB = db

	err = runMigrations(s.config, s.logger)
	if err != nil {
		s.logger.WithError(err).Error("Failed to Run Migrations!")
	}

	s.addrouter()
}

func (s *Server) Listen() {

	addr := fmt.Sprintf("%s:%d", s.ADDRESS, s.PORT)

	s.logger.Infof("Starting Server On: %s", addr)

	err := http.ListenAndServe(addr, s.Router)
	if err != nil {
		panic(err)
	}

}
