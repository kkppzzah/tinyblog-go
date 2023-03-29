// Package service 定义各项服务。
package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
	"ppzzl.com/tinyblog-go/recommend/common"
	"ppzzl.com/tinyblog-go/recommend/interfaces"
	"ppzzl.com/tinyblog-go/recommend/model"
)

// ArticleEventHandlerKafka 文章事件处理。
type ArticleEventHandlerKafka struct {
	kafkaReader       *kafka.Reader
	articleRepository interfaces.ArticleRepository
}

// NewArticleEventHandlerKafka 创建ArticleEventHandlerKafka实例。
func NewArticleEventHandlerKafka(ctx interfaces.Context) *ArticleEventHandlerKafka {
	brokers := strings.Split(common.MustGetEnv(common.EnvKafkaBrokers, ""), ",")
	h := &ArticleEventHandlerKafka{
		kafkaReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			GroupID: common.KafkaTopicNameArticleEventConsumerGroup,
			Topic:   common.KafkaTopicNameArticleEvent,
		}),
		articleRepository: ctx.GetArticleRepository(),
	}
	go h.run()
	return h
}

func (h *ArticleEventHandlerKafka) run() {
	ctx := context.Background()
	for {
		m, err := h.kafkaReader.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := h.handleEvent(m.Value); err != nil {
			log.Printf("failed to handle user event message, %s, %v", string(m.Value), err)
			continue
		}
		if err := h.kafkaReader.CommitMessages(ctx, m); err != nil {
			log.Printf("failed to commit messages: %v", err)
		}
	}
}

func (h *ArticleEventHandlerKafka) handleEvent(jsonData []byte) error {
	event := &interfaces.ArticleEvent{}
	if err := event.FromJSON(jsonData); err != nil {
		return err
	}
	switch event.EventType {
	case interfaces.ArticleEventTypeCreate:
		return h.createUser(event)
	case interfaces.ArticleEventTypeUpdate:
		return h.updateUser(event)
	case interfaces.ArticleEventTypeDelete:
		return h.deleteUser(event)
	}
	return nil
}

func (h *ArticleEventHandlerKafka) createUser(event *interfaces.ArticleEvent) error {
	article := &model.Article{
		ID: event.ArticleInfo.ID,
	}
	_, err := h.articleRepository.Create(article)
	return err
}

func (h *ArticleEventHandlerKafka) updateUser(userEvent *interfaces.ArticleEvent) error {
	return nil
}

func (h *ArticleEventHandlerKafka) deleteUser(userEvent *interfaces.ArticleEvent) error {
	// TODO
	return nil
}
