// Package service 定义各项服务。
package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
	"ppzzl.com/tinyblog-go/article/common"
	"ppzzl.com/tinyblog-go/article/interfaces"
	"ppzzl.com/tinyblog-go/article/model"
)

// UserEventHandlerKafka 用户事件处理。
type UserEventHandlerKafka struct {
	kafkaReader    *kafka.Reader
	userRepository interfaces.UserRepository
}

// NewUserEventHandlerKafka 创建UserEventHandlerKafka实例。
func NewUserEventHandlerKafka(ctx interfaces.Context) *UserEventHandlerKafka {
	brokers := strings.Split(common.MustGetEnv(common.EnvKafkaBrokers, ""), ",")
	h := &UserEventHandlerKafka{
		kafkaReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			GroupID: common.KafkaTopicNameUserEventConsumerGroup,
			Topic:   common.KafkaTopicNameUserEvent,
		}),
		userRepository: ctx.GetUserRepository(),
	}
	go h.run()
	return h
}

func (h *UserEventHandlerKafka) run() {
	ctx := context.Background()
	for {
		m, err := h.kafkaReader.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := h.handleUserEvent(m.Value); err != nil {
			log.Printf("failed to handle user event message, %s, %v", string(m.Value), err)
			continue
		}
		if err := h.kafkaReader.CommitMessages(ctx, m); err != nil {
			log.Printf("failed to commit messages: %v", err)
		}
	}
}

func (h *UserEventHandlerKafka) handleUserEvent(jsonData []byte) error {
	userEvent := &interfaces.UserEvent{}
	if err := userEvent.FromJSON(jsonData); err != nil {
		return err
	}
	switch userEvent.EventType {
	case interfaces.UserEventTypeCreate:
		return h.createUser(userEvent)
	case interfaces.UserEventTypeUpdate:
		return h.updateUser(userEvent)
	case interfaces.UserEventTypeDelete:
		return h.deleteUser(userEvent)
	}
	return nil
}

func (h *UserEventHandlerKafka) createUser(userEvent *interfaces.UserEvent) error {
	user := &model.User{
		ID:       userEvent.UserInfo.ID,
		Name:     userEvent.UserInfo.Name,
		Nickname: userEvent.UserInfo.Nickname,
	}
	_, err := h.userRepository.Create(user)
	return err
}

func (h *UserEventHandlerKafka) updateUser(userEvent *interfaces.UserEvent) error {
	user := &model.User{
		ID:       userEvent.UserInfo.ID,
		Name:     userEvent.UserInfo.Name,
		Nickname: userEvent.UserInfo.Nickname,
	}
	err := h.userRepository.Update(user)
	return err
}

func (h *UserEventHandlerKafka) deleteUser(userEvent *interfaces.UserEvent) error {
	// TODO
	return nil
}
