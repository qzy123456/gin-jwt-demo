package redis

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

// 上锁
// lockKey 锁的名称;
// ex 超时时间;
// retry 重试次数 >= 1;
func (p *Pool) Lock(lockKey string, ex uint, retry int) error {
	conn := p.Conn.Get()
	defer conn.Close()

	ts := time.Now() // as random value
	for i := 1; i <= retry; i++ {
		if i > 1 { // sleep if not first time
			time.Sleep(time.Second)
		}
		v, err := conn.Do("SET", lockKey, ts, "EX", ex, "NX")
		if err == nil {
			if v == nil {
				log.Printf("get %s lock failed, retry times:%d", lockKey, i)
			} else {
				log.Printf("get %s lock success", lockKey)
				break
			}
		} else {
			log.Printf("get %s lock failed with err: %s \n", lockKey, err)
		}
		if i >= retry {
			err = fmt.Errorf("get lock failed with max retry times")
			return err
		}
	}
	return nil
}

// 解锁
func (p *Pool) UnLock(lockKey string) error {
	conn := p.Conn.Get()
	defer conn.Close()

	v, err := redis.Bool(conn.Do("DEL", lockKey))
	if err == nil {
		if v {
			return nil
		} else {
			return errors.New("unLock failed")
		}
	} else {
		return errors.New("unLock failed, err:" + err.Error())
	}
}
