package repositories

import (
	"context"
	"errors"
	"github.com/dipay/internal"
	"github.com/dipay/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type companyRepository struct {
	MongoDatabase  *mongo.Database
	CompaniesModel model.ICompanies
}

type ICompanyRepository interface {
	FetchOne(ctx context.Context, filter interface{}, result interface{}) error
	Fetch(ctx context.Context, filter interface{}) ([]*model.Companies, error)
	Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error)
	Update(ctx context.Context, filter interface{}, update interface{}) error
}

func NewCompanyRepository(mongoDatabase *mongo.Database, companiesModel model.ICompanies) ICompanyRepository {
	return &companyRepository{
		MongoDatabase:  mongoDatabase,
		CompaniesModel: companiesModel,
	}
}

func (c *companyRepository) FetchOne(ctx context.Context, filter interface{}, result interface{}) error {
	companyTable := c.CompaniesModel.GetTableName()
	collection := c.MongoDatabase.Collection(companyTable)
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (c *companyRepository) Fetch(ctx context.Context, filter interface{}) ([]*model.Companies, error) {
	companyTable := c.CompaniesModel.GetTableName()
	collection := c.MongoDatabase.Collection(companyTable)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var companies []*model.Companies
	for cur.Next(context.TODO()) {
		var company = &model.Companies{}
		err := cur.Decode(&company)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, err
}

func (c *companyRepository) Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error) {
	companyTable := c.CompaniesModel.GetTableName()
	collection := c.MongoDatabase.Collection(companyTable)
	res, err := collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	lastInsertID := res.InsertedID.(primitive.ObjectID)
	return &lastInsertID, nil
}

func (c *companyRepository) Update(ctx context.Context, filter interface{}, update interface{}) error {
	companyTable := c.CompaniesModel.GetTableName()
	collection := c.MongoDatabase.Collection(companyTable)
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return errors.New(internal.ErrNoModifyUpdate.String())
	}

	return nil
}
