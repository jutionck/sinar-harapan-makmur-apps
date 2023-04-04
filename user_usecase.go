package main

type UserUseCase interface {
	RegisterNewUser(payload UserCredential) error
	FindAll() ([]UserCredential, error)
}

type userUseCase struct {
	repo UserRepository
}

func (u *userUseCase) RegisterNewUser(payload UserCredential) error {
	return u.repo.Create(payload)
}

func (u *userUseCase) FindAll() ([]UserCredential, error) {
	return u.repo.List()
}

func NewUserUseCase(repo UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
