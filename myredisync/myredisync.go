package myredisync

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redsync/redsync/v4"                  // 引入 redsync 库，用于实现基于 Redis 的分布式锁
	"github.com/go-redsync/redsync/v4/redis/goredis/v9" // 引入 redsync 的 goredis 连接池
	goredislib "github.com/redis/go-redis/v9"
)

func RedisSync(src string) {
	client := goredislib.NewClient(&goredislib.Options{
		Addr:     src,
		Password: "",
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex("redis-sync", redsync.WithExpiry(5*time.Second))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := mutex.LockContext(ctx); err != nil {
		panic(err)
	}

	strCh := make(chan struct{})
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	go func() {
		for {
			select {
			case <-ticker.C:
				if ok, err := mutex.ExtendContext(ctx); err != nil {
					panic(err)
				} else if ok {
					log.Println("redis sync extend lock")
				}
			case <-strCh:
				log.Println("redis sync unlock")
				return
			}
		}
	}()
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("redis sync lock %ds\n", i+1)
	}
	strCh <- struct{}{}

	if _, err := mutex.UnlockContext(ctx); err != nil {
		panic(err)
	}

}
