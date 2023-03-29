// Package service 定义各个service。
package service

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"ppzzl.com/tinyblog-go/article/common"
	"ppzzl.com/tinyblog-go/article/interfaces"
)

// ArticleEventPublisherKafka 发布文章事件的服务。
type ArticleEventPublisherKafka struct {
	kafkaWriter *kafka.Writer
	eventChan   chan *interfaces.ArticleEvent
	context     context.Context
	running     bool
}

// NewArticleEventPublisherKafka 创建ArticleEventPublisherKafka实例。
func NewArticleEventPublisherKafka() *ArticleEventPublisherKafka {
	brokers := strings.Split(common.MustGetEnv(common.EnvKafkaBrokers, ""), ",")
	p := &ArticleEventPublisherKafka{
		kafkaWriter: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    common.KafkaTopicNameArticleEvent,
			Balancer: &kafka.LeastBytes{},
		},
		eventChan: make(chan *interfaces.ArticleEvent),
		context:   context.Background(),
		running:   true,
	}
	go p.run()
	return p
}

// Publish 发布文章相关的事件。返回nil，不代表真正发送功能了；暂不考虑失败场景。
func (p *ArticleEventPublisherKafka) Publish(articleEvent *interfaces.ArticleEvent) error {
	p.eventChan <- articleEvent
	return nil
}

func (p *ArticleEventPublisherKafka) run() {
	for {
		event := <-p.eventChan
		if event == nil {
			if p.running {
				log.Printf("failed to receive article event from channel, something may be wrong")
			}
			log.Printf("ArticleEventPublisherKafka exit ...")
			break
		}
		err := p.doPublish(event)
		if err != nil {
			log.Printf("failed to publish article event, article id: %d, event: %s, %v", event.ArticleInfo.ID, event.EventType, err)
		}
	}
}

func (p *ArticleEventPublisherKafka) doPublish(event *interfaces.ArticleEvent) error {
	jsonRep, err := event.ToJSON()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(p.context, 1000*time.Millisecond)
	defer cancel()
	err = p.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.FormatInt(event.ArticleInfo.ID, 10)),
		Value: jsonRep,
	})
	return err
}
