package student

import "gorm.io/gorm"

type Repo struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) *Repo {
	return &Repo{DB: DB}
}

func (r *Repo) Create(student *Student) error  {

	err := r.DB.Create(student).Error
	return err
}

func (r *Repo) Update(student *Student) error  {

	err := r.DB.Updates(student).Where("roll = ?",student.Roll).Error
	return err
}

func (r *Repo) Get(roll int) (*Student, error) {

	student := &Student{}
	err := r.DB.Take(student, "roll = ?", roll).Error
	return student, err
}

func (r *Repo) GetAll() ([]*Student, error) {
	
	student := make([]*Student,0)
	err := r.DB.Model(&Student{}).Find(&student).Error
	return student,err
	
}

func (r *Repo) Delete(student *Student) error  {

	err := r.DB.Delete(student).Where("roll = ?",student.Roll).Error
	return err
}
