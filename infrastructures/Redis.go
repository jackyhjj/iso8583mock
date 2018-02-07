// Package infrastructures redis connection
package infrastructures

import (
	"fmt"
	"time"

	redisClient "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// Redis is the redis client structure
type Redis struct {
	Host           string
	Port           string
	DB             int
	Password       string
	currentSession *redisClient.Client
	//prefix         string
}

var (
	rds *Redis
)

// init the redis session
func init() {
	rds = &Redis{}
}

// NewRedis initialize redis client connection
func NewRedis(prefix string) *Redis {
	rds := &Redis{
		Host:     config.GetString("redis.host"),
		Port:     config.GetString("redis.port"),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
		//prefix:   prefix,
	}

	log.Error("redis host:", config.GetString("redis.host"))

	rds.Open(prefix)

	return rds
}

// setConfig redis
func (r *Redis) setConfig(prefix string) {
	r.Host = config.GetString("redis.host")
	r.Port = config.GetString("redis.port")
	r.Password = config.GetString("redis.password")
	r.DB = config.GetInt("redis.db")
	//r.prefix = config.GetString("redis.prefix")
}

// SetPrefix key redis
/*func SetPrefix(prefix string) *Redis { rds.prefix = prefix; return rds }
func (r *Redis) SetPrefix(prefix string) *Redis {
	r.prefix = prefix
	return r
}*/

// GetPrefix key redis
/*func GetPrefix() string { return rds.GetPrefix() }
func (r *Redis) GetPrefix() string {
	return r.prefix
}*/

// Open redis connection
func (r *Redis) Open(prefix string) *Redis {
	r.setConfig(prefix)
	client := redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
	})
	r.currentSession = client
	return r
}

// GetSession current session connection
func GetSession(prefix string) *redisClient.Client { return rds.GetSession(prefix) }
func (r *Redis) GetSession(prefix string) *redisClient.Client {
	if r.currentSession == nil {
		r.Open(prefix)
	}
	if r.ping() != nil {
		r.Open(prefix)
	}
	return r.currentSession
}

// ping connection redis
func (r *Redis) ping() error {
	err := r.currentSession.Ping().Err()
	if err != nil {
		log.Errorf("Redis connection error: %v", err)
		return err
	}
	return nil
}

// GetCached redis
func GetCached(prefix string, key string) *redisClient.StringCmd { return rds.GetCached(prefix,key) }
func (r *Redis) GetCached(prefix string, key string) *redisClient.StringCmd {
	s := r.GetSession(prefix)
	//newKey := fmt.Sprintf("%s%s", r.GetPrefix(), key)
	newKey := fmt.Sprintf("%s%s", prefix, key)
	return s.Get(newKey)
}

// SetCached redis
func SetCached(prefix string,key string, v interface{}, expiration time.Duration) error {
	return rds.SetCached(prefix,key, v, expiration)
}
func (r *Redis) SetCached(prefix string,key string, v interface{}, expiration time.Duration) error {
	s := r.GetSession(prefix)
	//newKey := fmt.Sprintf("%s%s", r.GetPrefix(), key)
	newKey := fmt.Sprintf("%s%s", prefix, key)
	err := s.Set(newKey, v, expiration).Err()
	if err != nil {
		log.Errorf("Redis error: %v", err)
		return err
	}

	return nil
}
