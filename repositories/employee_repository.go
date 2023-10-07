package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dipay/internal"
	"github.com/dipay/internal/db"
	"github.com/dipay/model"
	"github.com/dipay/pkg/httpClient"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type employeeRepository struct {
	MongoDatabase db.Database
	EmployeeModel model.IEmployees
	clientHttp    httpClient.IClientHttp
	hostEmail     string
}

//go:generate mockery --name IEmployeeRepository
type IEmployeeRepository interface {
	Fetch(ctx context.Context, filter interface{}) ([]*model.Employees, error)
	FetchOne(ctx context.Context, filter interface{}, result interface{}) error
	Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error)
	Update(ctx context.Context, filter interface{}, update interface{}) error
	Delete(ctx context.Context, filter interface{}) error
	SendEmail(ctx context.Context, email string) error
}

func NewEmployeeRepository(mongoDatabase db.Database, employeeModel model.IEmployees, client httpClient.IClientHttp, hostEmail string) IEmployeeRepository {
	return &employeeRepository{
		MongoDatabase: mongoDatabase,
		EmployeeModel: employeeModel,
		clientHttp:    client,
		hostEmail:     hostEmail,
	}
}

func (e *employeeRepository) Fetch(ctx context.Context, filter interface{}) ([]*model.Employees, error) {
	companyTable := e.EmployeeModel.GetTableName()
	collection := e.MongoDatabase.Collection(companyTable)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var employees []*model.Employees
	for cur.Next(context.TODO()) {
		var employee = &model.Employees{}
		err := cur.Decode(&employee)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, err
}

func (e *employeeRepository) FetchOne(ctx context.Context, filter interface{}, result interface{}) error {
	userAdminTable := e.EmployeeModel.GetTableName()
	collection := e.MongoDatabase.Collection(userAdminTable)
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) Create(ctx context.Context, model interface{}) (*primitive.ObjectID, error) {
	companyTable := e.EmployeeModel.GetTableName()
	collection := e.MongoDatabase.Collection(companyTable)
	resLastInsertedID, err := collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	lastInsertID := resLastInsertedID.(primitive.ObjectID)
	return &lastInsertID, nil
}

func (e *employeeRepository) Update(ctx context.Context, filter interface{}, update interface{}) error {
	employeeTable := e.EmployeeModel.GetTableName()
	collection := e.MongoDatabase.Collection(employeeTable)
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return errors.New(internal.ErrNoModifyUpdate.String())
	}

	return nil
}

func (e *employeeRepository) Delete(ctx context.Context, filter interface{}) error {
	employeeTable := e.EmployeeModel.GetTableName()
	collection := e.MongoDatabase.Collection(employeeTable)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeRepository) SendEmail(ctx context.Context, email string) error {
	data := map[string]string{
		"email": email,
	}
	fmt.Println("send email", email)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = e.clientHttp.Post(e.hostEmail, jsonData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
