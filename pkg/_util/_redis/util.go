package _redis

import (
	"adminbg/pkg/_util/_typ/_str"
	"github.com/go-redis/redis"
)

func IsNilErr(err error) bool {
	return err == nil || err == redis.Nil
}

func IsRedisErr(err error) bool {
	return !IsNilErr(err)
}

func StringMap(s []string) map[string]string {
	var m = make(map[string]string, len(s)/2)
	var tmpK string

	_str.Each(s, func(i int, elem string) {
		if i%2 == 0 {
			tmpK = elem
			return
		}
		m[tmpK] = elem
	})
	return m
}
