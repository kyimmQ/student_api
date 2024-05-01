package models

type Student struct {
	STD_ID        int    `json:"stdId"`
	STD_NAME      string `json:"stdName"`
	ACADEMIC_YEAR int    `json:"academicYear"`
}
