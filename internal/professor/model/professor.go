package professormodel

type Professor struct {
	Prof_id int `json:"prof_id" gorm:"column:PROF_ID;primaryKey"` 
	Prof_lname string `json:"prof_name" gorm:"column:PROF_NAME"`
	Prof_fname string `json:"prof_fname" gorm:"column:PROF_FNAME"`
}

func (Professor) TableName() string { 
	return "PROFESSOR"
}