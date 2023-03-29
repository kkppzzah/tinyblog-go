// Package service 定义各项服务。
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/microcosm-cc/bluemonday"
	"github.com/segmentio/kafka-go"
	"ppzzl.com/tinyblog-go/search/common"
	pb "ppzzl.com/tinyblog-go/search/genproto/article"
	"ppzzl.com/tinyblog-go/search/interfaces"
	"ppzzl.com/tinyblog-go/search/model"
)

// ArticleEventHandlerKafka 文章事件处理。
type ArticleEventHandlerKafka struct {
	kafkaReader    *kafka.Reader
	esClient       *elasticsearch.Client
	articleService interfaces.ArticleService
	ctx            context.Context
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
		esClient:       ctx.GetEsClient(),
		articleService: ctx.GetArticleService(),
		ctx:            context.Background(),
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
		return h.index(event)
	case interfaces.ArticleEventTypeUpdate:
		return h.index(event)
	case interfaces.ArticleEventTypeDelete:
		return h.delete(event)
	}
	return nil
}

func (h *ArticleEventHandlerKafka) index(event *interfaces.ArticleEvent) error {
	articleID := event.ArticleInfo.ID
	// 获取文章详情。
	ctx, cancel := context.WithTimeout(h.ctx, time.Second*10)
	defer cancel()
	rsp, err := h.articleService.Get(ctx, &pb.GetRequest{
		Id: articleID,
	})
	if err != nil {
		log.Printf("failed to index article （failed to get article detail）, article id: %d, %v", articleID, err)
		return err
	}
	// 索引文章。
	data, err := json.Marshal(model.Article{
		ID:          articleID,
		UserID:      rsp.UserId,
		Title:       rsp.Title,
		Summary:     rsp.Summary,
		Content:     h.clearContent(rsp.Content),
		PublishTime: time.Unix(rsp.PublishTime, 0),
	})
	req := esapi.IndexRequest{
		Index:      common.EsIndexNameArticle,
		DocumentID: strconv.FormatInt(articleID, 10),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	ctxIndex, cancelIndex := context.WithTimeout(h.ctx, time.Second*10)
	defer cancelIndex()
	res, err := req.Do(ctxIndex, h.esClient)
	if err != nil {
		log.Printf("failed to index article, article id: %d, %v", articleID, err)
		return err
	}
	if res.IsError() {
		log.Printf("failed to index article, article id: %d", articleID)
	}
	return nil
}

func (h *ArticleEventHandlerKafka) delete(event *interfaces.ArticleEvent) error {
	// 删除文章。
	articleID := event.ArticleInfo.ID
	req := esapi.DeleteRequest{
		Index:      common.EsIndexNameArticle,
		DocumentID: strconv.FormatInt(articleID, 10),
	}
	ctxIndex, cancelIndex := context.WithTimeout(h.ctx, time.Second*10)
	defer cancelIndex()
	res, err := req.Do(ctxIndex, h.esClient)
	if err != nil {
		log.Printf("failed to delete article, article id: %d, %v", articleID, err)
		return err
	}
	if res.IsError() {
		log.Printf("failed to index article, article id: %d", articleID)
	}
	return nil
}

func (h *ArticleEventHandlerKafka) clearContent(content string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(content)
}
