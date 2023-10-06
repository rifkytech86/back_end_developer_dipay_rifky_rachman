package bootstrap

import (
	"context"
	"github.com/dipay/commons"
	"github.com/dipay/internal"
	"github.com/dipay/internal/db"
	"github.com/dipay/internal/env"
	"github.com/dipay/internal/jwt"
	"github.com/dipay/internal/validations"
	"time"
)

type Application struct {
	MongoDBClient db.Database
	ENV           *env.ENV
	Validator     validations.IValidator
	JWT           jwt.IJWTRSAToken
}

func NewInitializeBootstrap() Application {
	app := Application{}
	app.ENV = env.NewENV()
	app.MongoDBClient = initialMongoDB(app.ENV.DatabaseURL, app.ENV.DatabaseName)
	app.Validator = validations.NewCustomValidator()
	registerValidatorCustom(app.Validator)

	// JWT
	app.JWT = initialJWT()

	return app
}

func initialMongoDB(databaseURL string, databaseName string) db.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//initialConnection := db.NewMongoDB(databaseURL, databaseName)

	mongoClient, err := db.NewClient(databaseURL)
	if err != nil {
		panic("error mongo connection")
	}

	err = mongoClient.Ping(ctx)
	if err != nil {
		panic("error mongo connection")
	}

	client := mongoClient.Database(databaseName)

	//err = mongoClient.Ping(context.TODO())
	//if err != nil {
	//	panic("error ping mongo")
	//}
	//
	//mongoClient, err := initialConnection.InitConnection()
	//if err != nil {
	//	panic("error mongo connection")
	//}
	//err = initialConnection.PingConnection(mongoClient)
	//if err != nil {
	//	panic("error mongo connection")
	//}
	//
	//handleMongoDB := initialConnection.SetDatabase(mongoClient, databaseName)

	return client
}

func registerValidatorCustom(validator validations.IValidator) {
	if err := validator.RegisterValidation(validations.ValidPassword, validations.ValidatorPassword); err != nil {
		panic(internal.ErrorInternalServer.String())
	}

	if err := validator.RegisterValidation(validations.ValidUsername, validations.ValidatorUsername); err != nil {
		panic(internal.ErrorInternalServer.String())
	}

	if err := validator.RegisterValidation(validations.ValidAddress, validations.ValidatorAddress); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
	if err := validator.RegisterValidation(validations.ValidCompanyName, validations.ValidatorCompanyName); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
	if err := validator.RegisterValidation(validations.ValidPhoneNumber, validations.ValidatePhoneNumber); err != nil {
		panic(internal.ErrorInternalServer.String())
	}

	if err := validator.RegisterValidation(validations.ValidEmail, validations.ValidateEmail); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
	if err := validator.RegisterValidation(validations.ValidJobTitle, validations.ValidateJobTitle); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
	if err := validator.RegisterValidation(validations.ValidEmployee, validations.ValidateEmployee); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
	if err := validator.RegisterValidation(validations.ValidDuplicateZero, validations.ValidateDuplicateZero); err != nil {
		panic(internal.ErrorInternalServer.String())
	}
}

func initialJWT() jwt.IJWTRSAToken {
	reader := commons.NewFileReader()
	privateKeyBytes, err := reader.ReadFile("private_key.pem")
	if err != nil {
		panic(internal.ErrorInternalReadFile.String())
	}
	publicKeyBytes, err := reader.ReadFile("public_key.pem")
	if err != nil {
		panic(internal.ErrorInternalReadFile.String())
	}

	return jwt.NewJWTRSAToken(privateKeyBytes, publicKeyBytes)
}
