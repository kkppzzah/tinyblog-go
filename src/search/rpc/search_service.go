// Package rpc 处理gRPC服务请求。
package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"ppzzl.com/tinyblog-go/search/common"
	pb "ppzzl.com/tinyblog-go/search/genproto/search"
	"ppzzl.com/tinyblog-go/search/interfaces"
)

// SearchService 实现推荐服务。
type SearchService struct {
	esClient *elasticsearch.Client
	ctx      context.Context
}

// NewSearchService 创建RecommendService实例。
func NewSearchService(ctx interfaces.Context) *SearchService {
	rs := &SearchService{
		esClient: ctx.GetEsClient(),
		ctx:      context.Background(),
	}
	return rs
}

// SimpleSearch 简单搜索。
func (svc *SearchService) SimpleSearch(ctx context.Context, req *pb.SimpleSearchRequest) (*pb.SimpleSearchResponse, error) {
	content := req.Content
	// 构造查询。
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  content,
				"fields": []string{"title", "content"},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("failed to encode search request, content: %s, %v", content, err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(svc.ctx, time.Second*10)
	defer cancel()
	// 执行查询。
	res, err := svc.esClient.Search(
		svc.esClient.Search.WithContext(ctx),
		svc.esClient.Search.WithIndex(common.EsIndexNameArticle),
		svc.esClient.Search.WithBody(&buf),
		svc.esClient.Search.WithTrackTotalHits(true),
		svc.esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Printf("failed to search, content: %s, %v", content, err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Printf("failed to parse the response body, content: %s, %v", content, err)
		} else {
			log.Printf("failed to parse the response body, content: %s, %s, %v", content, res.Status(), e["error"])
		}
		return nil, common.NewError(common.ErrorCodeInternalError, nil)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("failed to parse the response body, content: %s, %v", content, err)
		return nil, common.NewError(common.ErrorCodeInternalError, err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	artciles := make([]*pb.UserArticle, len(hits))
	rsp := &pb.SimpleSearchResponse{
		Articles: artciles,
	}
	for i, hit := range hits {
		s := hit.(map[string]interface{})["_source"].(map[string]interface{})
		artciles[i] = &pb.UserArticle{
			Id:     int64(s["id"].(float64)),
			UserId: int64(s["user_id"].(float64)),
			// TODO 暂时返回这些。
		}
	}
	return rsp, nil
}
