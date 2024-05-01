package main

import (
	"errors"
	"kyimmQ/student_api/internal/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) CreateStudent(w http.ResponseWriter, r *http.Request) {

	var student models.Student
	err := app.readJSON(w, r, &student)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	if student.STD_ID == 0 || student.STD_NAME == "" || student.ACADEMIC_YEAR == 0 {
		app.errorJSON(w, errors.New("missing required fields or invalid value in request body"), http.StatusBadRequest)
		return
	}
	result, statusCode, err := app.DB.CreateStudent(student.STD_ID, student.STD_NAME, student.ACADEMIC_YEAR)

	if !result {
		app.errorJSON(w, err, statusCode)
		return
	}

	_ = app.writeJSON(w, statusCode, &student)
}

func (app *application) GetStudentById(w http.ResponseWriter, r *http.Request) {

	sid, err := strconv.Atoi(chi.URLParam(r, "sid"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid id provided, try again"), http.StatusBadRequest)
		return
	}
	student, statusCode, err := app.DB.SearchStudentByID(sid)
	if err != nil {
		app.errorJSON(w, err, statusCode)
		return
	}
	_ = app.writeJSON(w, statusCode, student)
}

func (app *application) GetStudentByName(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	if !queryValues.Has("name") {
		app.errorJSON(w, errors.New("receive no query string, try again"), http.StatusBadRequest)
		return
	}
	name := queryValues.Get("name")
	students, statusCode, err := app.DB.SearchStudentByName(name)
	if err != nil {
		app.errorJSON(w, err, statusCode)
		return
	}
	_ = app.writeJSON(w, statusCode, students)
}
