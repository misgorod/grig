package repository

import "grig/internal/model"

func (r *Repository) GetGroupByID(id int) (model.StudyGroup, error) {
	var group model.StudyGroup
	err := r.db.Get(&group, `SELECT * FROM study_group WHERE study_group.id=$1`, id)
	return group, err
}

func (r *Repository) GetGroups() ([]model.StudyGroup, error) {
	var groups []model.StudyGroup
	err := r.db.Select(groups, `SELECT * FROM study_group`)
	return groups, err
}

func (r *Repository) UpdateGroup(group model.StudyGroup) error {
	_, err := r.db.Exec(`UPDATE study_group SET study_group.name=$2 WHERE study_group.id=$1`, group.ID, group.Name)
	return err
}

func (r *Repository) CreateGroup(group model.StudyGroup) error {
	_, err := r.db.Exec(`INSERT INTO study_group VALUES ($1)`, group.Name)
	return err
}

func (r *Repository) DeleteGroup(group model.StudyGroup) error {
	_, err := r.db.Exec(`DELETE study_group WHERE study_group.id=$1`, group.ID)
	return err
}
