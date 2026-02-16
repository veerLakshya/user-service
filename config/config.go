package config

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type JWT struct {
	SecretKey string
}

type DBConfig struct {
	URI      string
	HOST     string
	USERNAME string
	PASSWORD string
	NAME     string
	PORT     string
}

type ServerConfig struct {
	IdleTimeout  int
	ReadTimeout  int
	WriteTimeout int
	ShutdownWait int
	Port         int
}

type Configurations struct {
	DB     DBConfig
	SERVER ServerConfig
	JWT    JWT

	ENV     string
	BaseURL string
}

// GetAppConfig ...
func GetAppConfig() *Configurations {
	appConfig := LoadConfig()
	return &appConfig
}

// GetDBConn ...
func GetDBConn(ctx context.Context, log *zap.SugaredLogger, conf DBConfig) *mongo.Database {
	dbConn := configureDatabase(ctx, log, conf)
	return dbConn
}

// LoadConfig ...
func LoadConfig() Configurations {
	var env map[string]string
	var conf Configurations
	viper.SetConfigName(".env") // config file name
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // read value ENV variable
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	conf.DB = DBConfig{
		URI:  env["mongodb_url"],
		NAME: env["db_name"],
	}
	conf.ENV = env["env"]
	// Server Config
	port, _ := strconv.Atoi(env["server_port"])
	idleTimeout, _ := strconv.Atoi(env["server_idletimeout"])
	readTimeout, _ := strconv.Atoi(env["server_readtimeout"])
	shutdownTimeout, _ := strconv.Atoi(env["server_shutdownwait"])
	writeTimeout, _ := strconv.Atoi(env["server_writetimeout"])
	conf.SERVER = ServerConfig{
		Port:         port,
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		ShutdownWait: shutdownTimeout,
		WriteTimeout: writeTimeout,
	}

	conf.JWT = JWT{SecretKey: env["jwt_secret_key"]}
	conf.BaseURL = env["base_url"]
	return conf
}

func configureDatabase(ctx context.Context, lgr *zap.SugaredLogger, conf DBConfig) *mongo.Database {
	if conf.URI == "" {
		lgr.Fatal("Set MongoDB URI in your config.yaml file")
	}

	clientOpt := options.Client().ApplyURI(conf.URI)
	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		lgr.Fatal("Error occured connecting to the Database")
	}

	database := client.Database(conf.NAME)
	return database
}
