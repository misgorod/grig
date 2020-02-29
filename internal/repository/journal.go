package repository

import "grig/internal/model"

func (r *Repository) GetJournals() ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT * FROM journal
		JOIN student ON journal.student_id=student.id
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
	`
	err := r.db.Select(journals, query)
	return journals, err
}

func (r *Repository) GetJournalByStudent(id int) ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT * FROM journal
		JOIN student ON journal.student_id=student.id AND student.id=$1
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
	`
	err := r.db.Select(journals, query, id)
	return journals, err
}

func (r *Repository) GetJournalByStudyGroup(id int) ([]model.Journal, error) {
	var journals []model.Journal
	query := `
		SELECT * FROM journal
		JOIN student ON journal.student_id=student.id AND student.study_group_id=$1
		JOIN study_plan ON journal.study_plan_id=study_plan.id
		JOIN mark ON journal.mark_id=mark.id
	`
	err := r.db.Select(journals, query, id)
	return journals, err
}

func (r *Repository) CreateJournal(journal model.Journal) error {
	query := `INSERT INTO journal VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, journal.Student.ID, journal.StudyPlan.ID, journal.InTime, journal.Count, journal.Mark.ID)
	return err
}
