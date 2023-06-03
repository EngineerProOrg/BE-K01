package service

func (s *studentManagerService) GetStudent(id int64) interface{} {
	return s.repo.GetStudentByIdx(1)
}
