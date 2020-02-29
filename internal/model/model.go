package model

type Student struct {
	ID int `json:"studentId"`
	Name string `json:"studentName"`
	Surname string `json:"studentSurname"`
	SecondName string `json:"studentSecondName"`
	StudyGroup
}

type StudyGroup struct {
	ID int `json:"groupId"`
	Name string `json:"groupName"`
}

type Subject struct {
	ID int `json:"subjectId"`
	Name string `json:"subjectName"`
	ShortName string `json:"subjectShortName"`
}

type ExamType struct {
	ID int `json:"-"`
	Type string `json:"examType"`
}

type StudyPlan struct {
	ID int `json:"-"`
	Subject
	ExamType
}

type Mark struct {
	ID int `json:"-"`
	Name string `json:"MarkName"`
	Value string `json:"MarkValue"`
}

type Journal struct {
	ID int `json:"journalId,omitempty"`
	Student
	StudyPlan
	InTime bool `json:"journalInTime"`
	Count int `json:"journalCount"`
	Mark
}
