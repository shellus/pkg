package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/my/repo/queue"
	"time"
)

var redis_prefix = "queue"

type redisDrive struct {
	channelName string
	client      *redis.Client
	C           chan *queue.Job
	jobKey      string
	jobRollKey  string
	ctx         context.Context
	cancel      context.CancelFunc
}

func New(channelName string) (*redisDrive, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.20:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	if channelName == "" {
		channelName = "default"
	}

	ctx, cancel := context.WithCancel(context.Background())
	r := &redisDrive{
		channelName: channelName,
		client:      client,
		C:           make(chan *queue.Job),
		jobKey:      redis_prefix + ":" + channelName,
		jobRollKey:  redis_prefix + ":" + channelName + ":" + "roll",
		cancel:      cancel,
		ctx:         ctx,
	}

	go r.makeCh()

	return r, nil
}

func (r *redisDrive) Push(j *queue.Job) {
	r.client.LPush(r.jobKey, serialization(j))
}

func (r *redisDrive) Close() {
	r.cancel()
}
func (r *redisDrive) makeCh() {

	for {
		// 检查是否cancel
		if r.ctx.Err() != nil {
			fmt.Println("工作退出！")
			close(r.C)
			// 通知退出 err = "context canceled"
			return
		}

		// 从redis阻塞取数据，每秒跳出看看ctx有没有cancel
		line := r.client.BRPopLPush(r.jobKey, r.jobRollKey, time.Second)
		err := line.Err()

		// 如果没有数据就继续循环
		if err != nil && err.Error() == "redis: nil" {
			continue
		}

		// 如果有错误
		if err != nil && err.Error() != "redis: nil" {
			panic(fmt.Errorf("Redis Err: %s", err.Error()))
		}

		// 接收到redis数据
		s, err := line.Result()
		if err != nil {
			panic(err)
		}
		j := &queue.Job{}
		err = deserialization(s, j)
		if err != nil {
			panic(err)
		}
		r.C <- j
	}
}

func serialization(j *queue.Job) (s string) {
	buf, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(buf)
}

func deserialization(s string, j *queue.Job) (err error) {
	err = json.Unmarshal([]byte(s), j)
	if err != nil {
		return err
	}
	return nil
}
