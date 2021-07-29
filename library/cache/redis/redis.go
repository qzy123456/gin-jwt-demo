package redis

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func (p *Pool) Ping() (err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	_, err = redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return
}

// Set a key/value
func (p *Pool) Set(key string, data interface{}, time int) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	var value []byte
	value, err = json.Marshal(data)
	if err != nil {
		return
	}

	if time > 0 {
		reply, err = conn.Do("SETEX", key, time, value)
	} else {
		reply, err = conn.Do("SET", key, value)
	}
	return
}

// SetString Set a key/value
func (p *Pool) SetString(key string, str string, time int) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	if time > 0 {
		reply, err = conn.Do("SETEX", key, time, str)
	} else {
		reply, err = conn.Do("SET", key, str)
	}
	return
}

// Expire key timeout
func (p *Pool) Expire(key string, time int) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()
	reply, err = conn.Do("EXPIRE", key, time)
	return
}

// Exists check a key
func (p *Pool) Exists(key string) bool {
	conn := p.Conn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func (p *Pool) Get(key string) ([]byte, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	return reply, err
}

// GetString Get string
func (p *Pool) GetString(key string) (string, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	return reply, err
}

// GetStringMap Get string
func (p *Pool) GetStringMap(key string) (map[string]string, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.StringMap(conn.Do("GET", key))
	return reply, err
}

// Delete delete a key
func (p *Pool) Delete(key string) (bool, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LPush lpush a key 头部添加
func (p *Pool) LPush(key string, data interface{}) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	var value []byte
	value, err = json.Marshal(data)
	if err != nil {
		return
	}

	reply, err = conn.Do("LPUSH", key, value)
	return
}

// RPush rpush a key 尾部添加
func (p *Pool) RPush(key string, data interface{}) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	var value []byte
	value, err = json.Marshal(data)
	if err != nil {
		return
	}

	reply, err = conn.Do("RPUSH", key, value)
	return
}

// LPop lpop a key
func (p *Pool) LPop(key string) ([]byte, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("LPOP", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// RPop rpop a key
func (p *Pool) RPop(key string) ([]byte, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("RPOP", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// LikeDeletes batch delete
func (p *Pool) LikeDeletes(key string) error {
	conn := p.Conn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = p.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

// ZADD 向有序集合添加一个或多个成员，或者更新已存在成员的分数
func (p *Pool) ZADD(key string, id string, score int64) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err = conn.Do("ZADD", key, score, id)
	return
}

// ZRANGEBYSCORE 通过分数返回有序集合指定区间内的成员
func (p *Pool) ZRANGEWITHSCORES(key string, start, end interface{}) (reply map[string]int64, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err = redis.Int64Map(conn.Do("ZRANGE", key, start, end, "WITHSCORES"))
	return
}

// ZADD 向有序集合对指定成员的分数加上增量 incr
func (p *Pool) ZINCRBY(key string, id string, incr int64) (reply interface{}, err error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err = conn.Do("ZINCRBY", key, incr, id)
	return
}

// ZREVRANK a key
func (p *Pool) ZREVRANK(key string, id string) (int64, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.Int64(conn.Do("ZREVRANK", key, id))
	if err != nil {
		return 0, err
	}

	return reply+1, nil
}

// ZSCORE a key
func (p *Pool) ZSCORE(key string, id string) (string, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("ZSCORE", key, id))
	if err != nil {
		return "0", err
	}

	return reply, nil
}

// LPop incr a key
func (p *Pool) Incr(key string) (int64, error) {
	conn := p.Conn.Get()
	defer conn.Close()

	reply, err := redis.Int64(conn.Do("INCR", key))
	if err != nil {
		return 0, err
	}

	return reply, nil
}