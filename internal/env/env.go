package env

import (
	"github.com/spf13/viper"
	"log"
)

type ENV struct {
	DatabaseURL         string `mapstructure:"DATABASE_URL"`
	DatabaseName        string `mapstructure:"DATABASE_NAME"`
	DBMaxOpenConnection int    `mapstructure:"DB_MAX_OPEN_CONNECTION"`
	DBMaxIdleConnection int    `mapstructure:"DB_MAX_IDLE_CONNECTION"`
	ContextTimeOut      int    `mapstructure:"CONTEXT_TIMEOUT"`
	ExpiredAuthTime     int    `mapstructure:"EXPIRED_AUTH_TIME"`
	ExAPICountries      string `mapstructure:"EXTERNAL_API_COUNTRIES"`
	HTTPAddress         string `mapstructure:"HTTP_ADDRESS"`
}

func NewENV() *ENV {
	env := ENV{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

//env := ENV{}
//env.DatabaseURL = os.Getenv("DATABASE_URL")
//
//contextTimeOutStr := os.Getenv("CONTEXT_TIMEOUT")
//contextTime, err := strconv.Atoi(contextTimeOutStr)
//if err != nil {
//	contextTime = 20
//}
//env.ContextTimeOut = contextTime
//
//maxOpenStr := os.Getenv("MAX_OPEN_CONNECTION")
//maxOpenConnections, err := strconv.Atoi(maxOpenStr)
//if err != nil {
//	maxOpenConnections = 10
//}
//env.MaxOpenConnection = maxOpenConnections
//
//maxIdleConnectionsStr := os.Getenv("MAX_IDLE_CONNECTION")
//maxIdleConn, err := strconv.Atoi(maxIdleConnectionsStr)
//if err != nil {
//	maxIdleConn = 5
//}
//env.MaxIdleConnection = maxIdleConn
//
//expAuthTimeStr := os.Getenv("EXPIRED_AUTH_TIME")
//expAuthTime, err := strconv.Atoi(expAuthTimeStr)
//if err != nil {
//	expAuthTime = 1
//}
//env.ExpiredAuthTime = expAuthTime
//return &env
