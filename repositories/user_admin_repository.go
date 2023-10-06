package repositories

import (
	"context"
	"github.com/dipay/internal/db"
	"github.com/dipay/model"
)

type userAdminRepository struct {
	MongoDatabase  db.Database
	UserAdminModel model.IUserAdmin
}

type IUserAdminRepository interface {
	Fetch(ctx context.Context, filter interface{}, result interface{}) error
}

func NewUserAdminRepository(mongoDatabase db.Database, userAdminModel model.IUserAdmin) IUserAdminRepository {
	return &userAdminRepository{
		MongoDatabase:  mongoDatabase,
		UserAdminModel: userAdminModel,
	}
}

func (u *userAdminRepository) Fetch(ctx context.Context, filter interface{}, result interface{}) error {
	userAdminTable := u.UserAdminModel.GetTableName()
	collection := u.MongoDatabase.Collection(userAdminTable)
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
