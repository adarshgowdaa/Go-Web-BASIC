package home

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"time"
)

type Handler struct {
	logger    *logrus.Logger
	templates *template.Template
}

func NewHandler(logger *logrus.Logger, templates *template.Template) *Handler {
	return &Handler{logger: logger, templates: templates}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"time": time.Now().Format(time.RFC822),
	}

	w.Header().Add("Content-Type", "text/html")

	err := h.templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		h.logger.WithError(err).Error("Error Loading HOME PAGE!")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *Handler) Submit(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		h.logger.WithError(err).Error("Error Parsing Form!")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")

	if name == "" {
		http.Error(w, "Enter Name!!!", http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"name": name,
	}

	w.Header().Add("Content-Type", "text/html")

	err = h.templates.ExecuteTemplate(w, "submit.html", data)
	if err != nil {
		h.logger.WithError(err).Error("Error Loading HOME PAGE!")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
