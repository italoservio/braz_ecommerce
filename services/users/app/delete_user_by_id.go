package app

import (
	"github.com/italoservio/braz_ecommerce/packages/database"
	"github.com/italoservio/braz_ecommerce/services/users/infra/storage"
)

type DeleteUserByIdInterface interface {
	Do(id string) error
}

type DeleteUserByIdImpl struct {
	crudRepository database.CrudRepositoryInterface
	userRepository storage.UserRepositoryInterface
}

func NewDeleteUserByIdImpl(
	cr database.CrudRepositoryInterface,
	ur storage.UserRepositoryInterface,
) *DeleteUserByIdImpl {
	return &DeleteUserByIdImpl{crudRepository: cr, userRepository: ur}
}

func (gu *DeleteUserByIdImpl) Do(id string) error {
	err := gu.crudRepository.DeleteById(database.UsersCollection, id)
	if err != nil {
		return err
	}

	return nil
}