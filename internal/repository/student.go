package repository

import (
	"grig/internal/model"
)

func (r *Repository) CreateStudent(student model.Student) (int64, error) {
	query := `INSERT INTO student(name, surname, second_name, study_group_id) VALUES ($1, $2, $3, $4)`
	res, err := r.db.Exec(query, student.Name, student.Surname, student.SecondName, student.StudyGroup.ID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) GetStudentByID(id int) (model.Student, error) {
	var student model.Student
	err := r.db.Get(&student, `SELECT * FROM student WHERE student.id=$1`, id)
	return student, err
}

func (r *Repository) GetStudentsByGroup(id int) ([]model.Student, error) {
	var students []model.Student
	err := r.db.Select(&students, `SELECT * FROM student WHERE student.study_group_id=$1`, id)
	return students, err
}

func (r *Repository) GetStudents() ([]model.Student, error) {
	var students []model.Student
	err := r.db.Select(students, `SELECT * FROM student`)
	return students, err
}

func (r *Repository) DeleteStudent(id int) error {
	_, err := r.db.Exec(`DELETE FROM student WHERE student.id=$1`, id)
	return err
}

func (r *Repository) UpsertStudent(student model.Student) error {
	query := `
		INSERT INTO student VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (student.id) DO 
		UPDATE SET student.name=$2, student.surname=$3, student.second_name=$4, student.study_group_id=$5
	`
	_, err := r.db.Exec(query, student.ID, student.Name, student.Surname, student.SecondName, student.StudyGroup.ID)
	return err
}
