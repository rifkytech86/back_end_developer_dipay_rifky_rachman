package repositories

import (
	"context"
	"github.com/dipay/internal/db"
	"github.com/dipay/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userAdminRepository struct {
	MongoDatabase  db.Database
	UserAdminModel model.IUserAdmin
}

//go:generate mockery --name IUserAdminRepository
type IUserAdminRepository interface {
	Fetch(ctx context.Context, filter interface{}, result interface{}) error
	Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error)
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

func (u *userAdminRepository) Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error) {
	companyTable := u.UserAdminModel.GetTableName()
	collection := u.MongoDatabase.Collection(companyTable)
	resLastInsertedID, err := collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	lastInsertID := resLastInsertedID.(primitive.ObjectID)
	return &lastInsertID, nil
}
