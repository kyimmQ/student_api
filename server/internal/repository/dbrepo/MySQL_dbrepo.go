package dbrepo

import (
	"database/sql"
	"errors"
	"kyimmQ/student_api/internal/models"
	"net/http"
)

type MySQLDBRepo struct {
	DB *sql.DB
}

func (m *MySQLDBRepo) Connect() *sql.DB {
	return m.DB
}

// database interaction for creating new student
func (m *MySQLDBRepo) CreateStudent(stdId int, stdName string, acaYear int) (bool, int, error) {
	var result int
	stmt, err := m.DB.Prepare("CALL create_new_student(?,?,?)")
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	err = stmt.QueryRow(stdId, stdName, acaYear).Scan(&result)
	if err != nil {
		return false, http.StatusInternalServerError, err
	}
	if result == 0 {
		return false, http.StatusBadRequest, errors.New("student id exists, try again")
	}
	return true, http.StatusCreated, nil
}

// database interaction for searching student
func (m *MySQLDBRepo) SearchStudentByID(stdID int) (*models.Student, int, error) {
	var student models.Student

	stmt, err := m.DB.Prepare("CALL search_student_by_id(?)")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(stdID).Scan(&student.STD_ID, &student.STD_NAME, &student.ACADEMIC_YEAR)
	if err != nil {
		return nil, http.StatusNotFound, errors.New("id not found, try again")
	}

	return &student, http.StatusOK, nil
}

func (m *MySQLDBRepo) SearchStudentByName(stdName string) (*[]models.Student, int, error) {
	var students []models.Student
	var found = false
	stmt, err := m.DB.Prepare("CALL search_student_by_name(?)")
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(stdName)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	for rows.Next() {
		found = true
		var student models.Student
		err = rows.Scan(&student.STD_ID, &student.STD_NAME, &student.ACADEMIC_YEAR)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		students = append(students, student)
	}

	if !found {
		return nil, http.StatusNotFound, errors.New("name not found, try again")
	}

	return &students, http.StatusOK, nil
}
