package repositories

import (
	"context"
	"github.com/dipay/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userAdminRepository struct {
	MongoDatabase  *mongo.Database
	UserAdminModel model.IUserAdmin
}

type IUserAdminRepository interface {
	Fetch(ctx context.Context, filter interface{}, result interface{}) error
}

func NewUserAdminRepository(mongoDatabase *mongo.Database, userAdminModel model.IUserAdmin) IUserAdminRepository {
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
