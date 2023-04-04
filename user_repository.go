package main

type UserRepository interface {
	Create(payload UserCredential) error
	List() ([]UserCredential, error)
}

type userRepository struct {
	db []UserCredential
}

func (u *userRepository) Create(payload UserCredential) error {
	u.db = append(u.db, payload)
	return nil
}

func (u *userRepository) List() ([]UserCredential, error) {
	return u.db, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
