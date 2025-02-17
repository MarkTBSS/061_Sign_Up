package usersUsecases

import (
	"github.com/MarkTBSS/061_Sign_Up/config"
	"github.com/MarkTBSS/061_Sign_Up/modules/users"
	"github.com/MarkTBSS/061_Sign_Up/modules/users/usersRepositories"
)

type IUsersUsecase interface {
	InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error)
}

type usersUsecase struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UsersUsecase(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}

func (u *usersUsecase) InsertCustomer(req *users.UserRegisterReq) (*users.UserPassport, error) {
	// Hashing a password
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}
	// Insert user
	result, err := u.usersRepository.InsertUser(req, false)
	if err != nil {
		return nil, err
	}
	return result, nil
}
