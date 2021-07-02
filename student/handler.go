package student

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	_ "strconv"
)

type Handler struct {
	logger *logrus.Logger
	repo *Repo
}

func NewHandler(logger *logrus.Logger, repo *Repo) *Handler {
	return &Handler{logger: logger, repo: repo}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	student := &Student{}
	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		http.Error(w, "Invalid JSON!", http.StatusUnprocessableEntity)
		return
	}

	err = h.repo.Create(student)

	if err != nil {
		h.logger.WithError(err).Error("Error In Creating Students!")
		http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created!"))
	
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	pathVariable := mux.Vars(r)
	rollstr := pathVariable["roll"]

	h.logger.Debugf("ROLL: %v",rollstr)



	roll , err := strconv.Atoi(rollstr)
	if err != nil {
		http.Error(w, "Invalid Input!", http.StatusBadRequest)
		return
	}

	student,err:= h.repo.Get(roll)

	if errors.Is(err,gorm.ErrRecordNotFound)	{

		http.Error(w,"Student Does Not Exist!", http.StatusNotFound)
		return
	}

	if err != nil {
		h.logger.WithError(err).Error("Error In Fetching Students!")
		http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(student)
	if err != nil {

	}

}


func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	students ,err := h.repo.GetAll()
	if err != nil {
		h.logger.WithError(err).Error("Error In Fetching Students!")
		http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(students)
	if err != nil {

	}

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	student := &Student{}
	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		http.Error(w, "Invalid JSON!", http.StatusUnprocessableEntity)
		return
	}

	err = h.repo.Delete(student)

	if errors.Is(err,gorm.ErrRecordNotFound)	{

		http.Error(w,"Student Does Not Exist!", http.StatusNotFound)
		return
	}
	if err != nil {
		h.logger.WithError(err).Error("Error In Deleting Students!")
		http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Student Deleted!"))


}


func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

	student := &Student{}
	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		http.Error(w, "Invalid JSON!", http.StatusUnprocessableEntity)
		return
	}

	err = h.repo.Update(student)

	if errors.Is(err,gorm.ErrRecordNotFound)	{

		http.Error(w,"Student Does Not Exist!", http.StatusNotFound)
		return
	}
	if err != nil {
		h.logger.WithError(err).Error("Error In Updating Students!")
		http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Student Updated!"))
}
