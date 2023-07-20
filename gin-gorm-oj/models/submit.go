package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `json:"identity"`
	ProblemIdentity string `json:"problem_identity"`
	UserIdentity    string `json:"user_identity"`
	Path            string `json:"path"`
	Status          int    `json:"status"`
}

func (table *Submit) Tablename() string {
	return "submit"
}
