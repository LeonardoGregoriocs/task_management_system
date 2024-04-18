package user

type Repository interface {
	Save() error
	GetAll() error
	GetUserByCPF() error
	UpdateUser() error
	DeleteUser() error
}
