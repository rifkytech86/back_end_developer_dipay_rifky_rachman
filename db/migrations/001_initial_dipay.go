package migrations

import (
	"context"
	"fmt"
	"github.com/dipay/internal/db"
	"github.com/dipay/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Migration func(ctx context.Context, db db.Database) error

var AllMigrations = []Migration{
	MigrationUserAdmin,
	MigrationEmployee,
	MigrationCompanies,
}

func MigrationUserAdmin(ctx context.Context, db db.Database) error {
	modelUserAdmin := model.NewUserAdmin()
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := db.Collection(modelUserAdmin.GetTableName()).CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	initialUser := bson.D{
		{Key: "username", Value: "dipayusername"},
		{Key: "password", Value: "$2a$10$qoahj5NG6GzK.acOk0jn6ugflySm4dW.vVsjLMYg1pZJnAx74KsiS"},
		{Key: "created_at", Value: time.Now()},
		{Key: "updated_at", Value: time.Now()},
	}
	collection := db.Collection(modelUserAdmin.GetTableName())
	_, err = collection.InsertOne(ctx, initialUser)
	if err != nil {
		fmt.Println("warning migration table already exist")
		return err
	}
	return nil
}

func MigrationEmployee(ctx context.Context, db db.Database) error {
	keys := bson.D{
		{Key: "company_id", Value: 1},
		{Key: "email", Value: 1},
	}

	indexOptions := options.Index().SetUnique(true)

	modelEmployee := model.NewEmployees()
	_, err := db.Collection(modelEmployee.GetTableName()).CreateOne(ctx, mongo.IndexModel{
		Keys:    keys,
		Options: indexOptions,
	})
	if err != nil {
		return err
	}

	employee := bson.M{
		"name":         "dipaynameemployee",
		"email":        "dipayemail@email.com",
		"phone_number": "+62345345345222",
		"jobtitle":     "manager",
		"company_id":   "65206e8ad279a2723f4b9a52",
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
	}

	collection := db.Collection(modelEmployee.GetTableName())
	_, err = collection.InsertOne(ctx, employee)
	if err != nil {
		return err
	}
	return nil
}

func MigrationCompanies(ctx context.Context, db db.Database) error {
	keys := bson.D{
		{Key: "telephone_number", Value: 1},
	}

	indexOptions := options.Index().SetUnique(true)

	modelCompanies := model.NewCompanies()
	_, err := db.Collection(modelCompanies.GetTableName()).CreateOne(ctx, mongo.IndexModel{
		Keys:    keys,
		Options: indexOptions,
	})
	if err != nil {
		return err
	}

	document := bson.M{
		"company_name":     "dipay",
		"telephone_number": "+62899239393944",
		"address":          "address",
		"is_active":        true,
		"created_at":       time.Date(2023, 10, 6, 21, 22, 28, 119000000, time.UTC),
		"updated_at":       time.Date(2023, 10, 7, 3, 19, 42, 190000000, time.UTC),
	}

	collection := db.Collection(modelCompanies.GetTableName())
	_, err = collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	return nil
}
