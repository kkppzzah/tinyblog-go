// Package service 定义各个service。
package service

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"ppzzl.com/tinyblog-go/user/common"
	"ppzzl.com/tinyblog-go/user/interfaces"
)

// UserEventPublisherKafka 发布用户事件的服务。
type UserEventPublisherKafka struct {
	kafkaWriter   *kafka.Writer
	userEventChan chan *interfaces.UserEvent
	context       context.Context
	running       bool
}

// NewUserEventPublisherKafka 创建UserEventPublisherKafka实例。
func NewUserEventPublisherKafka() *UserEventPublisherKafka {
	brokers := strings.Split(common.MustGetEnv(common.EnvKafkaBrokers, ""), ",")
	p := &UserEventPublisherKafka{
		kafkaWriter: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    common.KafkaTopicNameUserEvent,
			Balancer: &kafka.LeastBytes{},
		},
		userEventChan: make(chan *interfaces.UserEvent),
		context:       context.Background(),
		running:       true,
	}
	go p.run()
	return p
}

// Publish 发布用户相关的事件。返回nil，不代表真正发送功能了；暂不考虑失败场景。
func (p *UserEventPublisherKafka) Publish(userEvent *interfaces.UserEvent) error {
	p.userEventChan <- userEvent
	return nil
}

func (p *UserEventPublisherKafka) run() {
	for {
		userEvent := <-p.userEventChan
		if userEvent == nil {
			if p.running {
				log.Printf("failed to receive user event from channel, something may be wrong")
			}
			log.Printf("UserEventPublisherKafka exit ...")
			break
		}
		err := p.doPublish(userEvent)
		if err != nil {
			log.Printf("failed to publish user event, user id: %d, event: %s, %v", userEvent.UserInfo.ID, userEvent.EventType, err)
		}
	}
}

func (p *UserEventPublisherKafka) doPublish(userEvent *interfaces.UserEvent) error {
	jsonRep, err := userEvent.ToJSON()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(p.context, 1000*time.Millisecond)
	defer cancel()
	err = p.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.FormatInt(userEvent.UserInfo.ID, 10)),
		Value: jsonRep,
	})
	return err
}
