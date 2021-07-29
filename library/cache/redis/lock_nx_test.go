package redis

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func init() {
	config = getConfig()
	p = NewPool(config)
}

func TestLock(t *testing.T) {
	var wg sync.WaitGroup

	key := "lock_demo"

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			// getLock
			err := p.Lock(key, 10, 10)
			if err != nil {
				fmt.Println(fmt.Sprintf("worker[%d] get lock failed:%v", id, err))
				return
			}
			// sleep for random
			for j := 0; j < 5; j++ {
				time.Sleep(time.Second)
				fmt.Println(fmt.Sprintf("worker[%d] hold lock for %ds", id, j+1))
			}
			// unLock
			err = p.UnLock(key)
			if err != nil {
				fmt.Println(fmt.Sprintf("worker[%d] unlock failed:%v", id, err))
			}
			fmt.Println(fmt.Sprintf("worker[%d] done", id))
		}(i)
	}

	wg.Wait()
	fmt.Println("demo is done!")
}
