package userRepository

import "github.com/rodrigoRVSN/beeus-api/src/domain/entity"

func (r *UserRepository) FindAllUsers(user *entity.User) ([]entity.User, error) {
	users := []entity.User{}

	response := r.DB.Find(&users)

	if response.Error != nil {
		return nil, response.Error
	}

	return users, nil
}
