package redisutils

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type RedisConn struct {
	Conn *redis.Client
}

var redisClient *redis.Client

func New(addr string) (*RedisConn, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:       addr,
		PoolSize:   100,
		MaxRetries: 2,
		Password:   "",
		DB:         0,
	})

	ping, err := conn.Ping().Result()
	if err == nil && len(ping) > 0 {
		println("Connected to Redis")
	} else {
		println("Redis Connection Failed")
	}
	return &RedisConn{
		Conn: conn,
	}, err
}

func (r *RedisConn) GetValue(key string) (interface{}, error) {
	var deserializedValue interface{}
	serializedValue, err := r.Conn.Get(key).Result()
	json.Unmarshal([]byte(serializedValue), &deserializedValue)
	return deserializedValue, err
}

func (r *RedisConn) SetValue(key string, value interface{}) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := r.Conn.Set(key, string(serializedValue), 0).Err()
	return true, err
}

func (r *RedisConn) SetValueWithTTL(key string, value interface{}, ttl int) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := r.Conn.Set(key, string(serializedValue), time.Duration(ttl)*time.Second).Err()
	return true, err
}

func (r *RedisConn) RPush(key string, valueList []string) (bool, error) {
	err := r.Conn.RPush(key, valueList).Err()
	return true, err
}

func (r *RedisConn) RpushWithTTL(key string, valueList []string, ttl int) (bool, error) {
	err := r.Conn.RPush(key, valueList, ttl).Err()
	return true, err
}

func (r *RedisConn) LRange(key string) (bool, error) {
	err := r.Conn.LRange(key, 0, -1).Err()
	return true, err
}

func (r *RedisConn) ListLength(key string) int64 {
	return r.Conn.LLen(key).Val()
}

func (r *RedisConn) Publish(channel string, message string) {
	r.Conn.Publish(channel, message)
}

func (r *RedisConn) GetKeyListByPattern(pattern string) []string {
	return r.Conn.Keys(pattern).Val()
}

func (r *RedisConn) IncrementValue(key string) int64 {
	return r.Conn.Incr(key).Val()
}

func (r *RedisConn) DelKey(key string) error {
	return r.Conn.Del(key).Err()
}
