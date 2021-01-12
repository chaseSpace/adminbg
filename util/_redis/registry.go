package _redis

import (
	"adminbg/util"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

/*
Redis registry, app can directly call them to initialize redis client.
*/

var (
	lock      sync.Mutex
	DefClient *redis.Client
)

// init default pool on registry
func MustInitDef(opts *redis.Options) {
	lock.Lock()
	defer lock.Unlock()
	util.Must(DefClient == nil, fmt.Errorf("_redis: DefClient already exists"))

	DefClient = newClient(opts)
}

func MustInit(opts *redis.Options) *redis.Client {
	return newClient(opts)
}

func newClient(opts *redis.Options) *redis.Client {
	cli := redis.NewClient(opts)
	_, err := cli.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("_redis: %v", err))
	}
	return cli
}

func Close() error {
	lock.Lock()
	defer lock.Unlock()
	err := DefClient.Close()
	DefClient = nil
	return err
}
