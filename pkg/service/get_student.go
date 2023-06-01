package service

func (s *studentManagerService) GetStudent(id int64) interface{} {
	s.repo.GetStudentByIdx(1)
}
