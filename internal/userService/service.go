package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) DeleteUserByID(id int) error {
	return s.repo.DeleteUserByID(id)
}

func (s *UserService) UpdateUserByID(id int, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}
