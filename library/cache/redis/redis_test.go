package redis

import (
	"testing"
)

var (
	p      *Pool
	config *Config
)

func init() {
	config = getConfig()
	p = NewPool(config)
}

func getConfig() (c *Config) {
	c = &Config{
		Host:     "127.0.0.1:6379",
		Password: "",
		Db:       0, // db
		MaxIdle:  2, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
	}

	return
}

func TestRedis(t *testing.T) {
	testSet(t, p)
	testExpire(t, p)
	testGet(t, p)
	testDelete(t, p)
	testLPush(t, p)
	testRPush(t, p)
	testLPop(t, p)
	testRPop(t, p)
	testLikeDeletes(t, p)
	if err := p.Close(); err != nil {
		t.Errorf("redis: close error(%v)", err)
	}
}

func testSet(t *testing.T, p *Pool) {
	var (
		key   = "test"
		value = "test"
	)

	if reply, err := p.Set(key, value, 0); err != nil {
		t.Errorf("redis: conn.Do(SET, %s, %s) error(%v)", key, value, err)
	} else {
		t.Logf("redis: set status: %s", reply)
	}
}

func testExpire(t *testing.T, p *Pool) {
	var (
		key = "test"
		i   = 300
	)

	if reply, err := p.Expire(key, i); err != nil {
		t.Errorf("redis: conn.Do(Expire, %s, %v) error(%v)", key, i, err)
	} else {
		t.Logf("redis: Expire status: %v", reply)
	}
}

func testGet(t *testing.T, p *Pool) {
	var (
		key = "test"
	)

	if reply, err := p.Get(key); err != nil {
		t.Errorf("redis: conn.Do(Get, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: get status: %s", reply)
	}
}

func testDelete(t *testing.T, p *Pool) {
	var (
		key = "test"
	)

	if reply, err := p.Delete(key); err != nil {
		t.Errorf("redis: conn.Do(Delete, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: delete status: %v", reply)
	}
}

func testLPush(t *testing.T, p *Pool) {
	var (
		key   = "test1"
		value = "test"
	)

	if reply, err := p.LPush(key, value); err != nil {
		t.Errorf("redis: conn.Do(Lpush, %s, %s) error(%v)", key, value, err)
	} else {
		t.Logf("redis: lpush status: %v", reply)
	}
}

func testRPush(t *testing.T, p *Pool) {
	var (
		key   = "test1"
		value = "test"
	)

	if reply, err := p.RPush(key, value); err != nil {
		t.Errorf("redis: conn.Do(Rpush, %s, %s) error(%v)", key, value, err)
	} else {
		t.Logf("redis: Rpush status: %v", reply)
	}
}

func testLPop(t *testing.T, p *Pool) {
	var (
		key = "test1"
	)

	if reply, err := p.LPop(key); err != nil {
		t.Errorf("redis: conn.Do(LPop, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: LPop status: %v", reply)
	}
}

func testRPop(t *testing.T, p *Pool) {
	var (
		key = "test1"
	)

	if reply, err := p.RPop(key); err != nil {
		t.Errorf("redis: conn.Do(RPop, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: RPop status: %v", reply)
	}
}

func testLikeDeletes(t *testing.T, p *Pool) {
	var (
		key = "te"
	)

	if err := p.LikeDeletes(key); err != nil {
		t.Errorf("redis: conn.Do(DELETE, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: LikeDeletes status: ok")
	}
}

func TestPool_ZAD(t *testing.T) {
	key := "test_sort_set"
	group1 := "group_1"
	group2 := "group_2"
	score := int64(10)

	_, err := p.ZADD(key, group1, score)
	if err != nil {
		t.Fatalf("ZADD key %s fail: %v", key, err)
	}

	t.Log(p.ZADD(key, group2, score))
}

func TestPool_ZRANGE(t *testing.T) {
	key := "test_sort_set"

	//t.Log(p.ZRANGEWITHSCORES(key, 0, -1))
	_, err := p.ZRANGEWITHSCORES(key, 0, -1)
	if err != nil {
		t.Fatalf("ZRANGEWITHSCORES key %s fail: %v", key, err)
	}
}
