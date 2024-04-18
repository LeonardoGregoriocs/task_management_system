package repository

import "database/sql"

type UserRepository struct {
	Db *sql.DB
}

func (c *UserRepository) Save() error {
	return nil
}

func (c *UserRepository) GetAll() error {
	return nil
}

func (c *UserRepository) GetUserByCPF() error {
	return nil
}

func (c *UserRepository) UpdateUser() error {
	return nil
}

func (c *UserRepository) DeleteUser() error {
	return nil
}
