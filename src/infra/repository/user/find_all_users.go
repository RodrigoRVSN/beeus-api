package userRepository

import (
	userDTO "github.com/rodrigoRVSN/beeus-api/src/application/dto/user"
	"github.com/rodrigoRVSN/beeus-api/src/domain/entity"
)

func (r *UserRepository) FindAllUsers(sortBy string, sortOrder string) (*[]userDTO.UserWithoutPasswordDTO, error) {
	rankedUsers := []entity.User{}

	query := r.DB.Order(sortBy + " " + sortOrder).Find(&rankedUsers)
	if query.Error != nil {
		return nil, query.Error
	}

	userDTOs := make([]userDTO.UserWithoutPasswordDTO, len(rankedUsers))
	for i, user := range rankedUsers {
		userDTOs[i] = userDTO.MapUserWithoutPasswordToDTO(user)
	}

	return &userDTOs, nil
}
