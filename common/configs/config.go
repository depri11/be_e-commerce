package configs

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/depri11/be_e-commerce/connections"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ConfigurationDomain interface {
	Read() (res *Configuration, err error)
}

type Configuration struct {
	PortApi         int
	PostgreConfig   postgreConfig
	RedisClient     *redis.Client
	AppEnv          string
	JwtKey          string
	PortUserService int
	UserService     string
	TimeoutCtx      time.Duration
}

type postgreConfig struct {
	Host     string
	User     string
	Port     string
	Database string
	Password string
	GormConn *gorm.DB
}

func NewConfiguration() *Configuration {
	ctx := time.Duration(20) * time.Second
	return &Configuration{TimeoutCtx: ctx}
}

func (c *Configuration) Read() (res *Configuration, err error) {
	portApiStr := os.Getenv("PORT_API")
	portApi, err := strconv.Atoi(portApiStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.PortApi = portApi

	postgreHost := os.Getenv("POSTGRE_HOST")
	if postgreHost == "" {
		err = errors.New("invalid-postgre-host")
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.Host = postgreHost
	postgreUser := os.Getenv("POSTGRE_USER")
	if postgreUser == "" {
		err = errors.New("invalid-postgre-user")
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.User = postgreUser
	postgrePort := os.Getenv("POSTGRE_PORT")
	if postgrePort == "" {
		err = errors.New("invalid-postgre-port")
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.Port = postgrePort
	postgreDb := os.Getenv("POSTGRE_DATABASE")
	if postgreDb == "" {
		err = errors.New("invalid-postgre-database")
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.Database = postgreDb
	postgrePass := os.Getenv("POSTGRE_PASSWORD")
	if postgrePass == "" {
		err = errors.New("invalid-postgre-credential")
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.Password = postgrePass
	c.AppEnv = os.Getenv("APP_ENV")
	if c.AppEnv == "" {
		err = errors.New("invalid-app-env")
		log.Println(err)
		return nil, err
	}

	gormConn, err := connections.GormConnect(c.PostgreConfig.Host, c.PostgreConfig.User, c.PostgreConfig.Password, c.PostgreConfig.Database, c.PostgreConfig.Port, c.AppEnv)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.PostgreConfig.GormConn = gormConn

	serviceUserPortStr := os.Getenv("SERVICE_USER_PORT")
	serviceUserPort, err := strconv.Atoi(serviceUserPortStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.PortUserService = serviceUserPort

	serviceUser := os.Getenv("SERVICE_USER")
	if serviceUser == "" {
		err = errors.New("invalid-service-user")
		log.Println(err)
		return nil, err
	}
	c.UserService = serviceUser

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		err = errors.New("invalid-redis-host")
		log.Println(err)
		return nil, err
	}
	redisPortStr := os.Getenv("REDIS_PORT")
	if redisPortStr == "" {
		err = errors.New("invalid-postgre-port")
		log.Println(err)
		return nil, err
	}
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	redisConn, err := connections.RedisConn(redisHost, redisPort)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c.RedisClient = redisConn
	c.JwtKey = os.Getenv("JWT_KEY")

	return c, nil
}
