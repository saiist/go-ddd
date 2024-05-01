package users

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (u *UserService) Exists(user User) (bool, error) {
	found, err := u.UserRepository.FindByName(&user.UserName)
	if err != nil {
		return false, err
	}

	return found != nil, nil
}
