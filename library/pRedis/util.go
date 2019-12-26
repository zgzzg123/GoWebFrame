package pRedis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"errors"
)

func Ping() error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func Get(key string) ([]byte, error) {

	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}

func Set(key string, value []byte) error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return err
}

func Exists(key string) (bool, error) {

	conn := Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func Delete(key string) error {

	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func GetKeys(pattern string) ([]string, error) {

	conn := Pool.Get()
	defer conn.Close()

	iter := 0
	keys := []string{}
	for {
		arr, err := redis.Values(conn.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}

func LPop(channel string) ([]byte,error){
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("lpop", channel))
	if err != nil {
		return data, fmt.Errorf("error getting channel %s: %v", channel, err)
	}

	return data, err
}


func RPush(channel string, value string) ([]byte,error){
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("rpush", channel, value))
	if err != nil {
		return data, err
	}

	return data, err
}

func LPush(channel string, value []byte) ([]byte,error){
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("lpush", channel, value))
	if err != nil {
		return data, err
	}

	return data, err
}

func Incr(counterKey string) (int, error) {

	conn := Pool.Get()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", counterKey))
}

func Sub(channel string){
	conn := Pool.Get()
	defer conn.Close()

	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channel)
}

func Pub(channel string, content []byte) interface{} {
	conn := Pool.Get()
	defer conn.Close()

	result,err := conn.Do("PUBLISH",channel,content)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}

func CountChannel(channel string) (int, error) {
	conn := Pool.Get()
	defer conn.Close()

	lenQueue, err := conn.Do("llen", channel)
	if err != nil {
		return 0, err
	}

	count, ok := lenQueue.(int64)
	if !ok {
		return 0, errors.New("类型转换错误!")
	}

	return int(count), nil
}