package hello

import (
	context "context"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
)

// ServiceImpl 实现服务。
type ServiceImpl struct {
	articles map[int32]*Article
}

// Server 服务端运行所需的各个组件。
type Server struct {
	host        string
	port        int
	grpcServer  *grpc.Server
	serviceImpl SearchServiceServer
}

// NewServer 创建服务端实例。
func NewServer(host string, port int) *Server {
	s := Server{
		host:        host,
		port:        port,
		serviceImpl: ServiceImpl{articles: make(map[int32]*Article)},
	}
	s.grpcServer = grpc.NewServer()
	return &s
}

// Run 运行服务端。
func (s *Server) Run() {
	address := fmt.Sprintf("%s:%d", s.host, s.port)
	l, _ := net.Listen("tcp", address)
	RegisterSearchServiceServer(s.grpcServer, s.serviceImpl)
	s.grpcServer.Serve(l)
}

// Stop 停止服务端运行。
func (s *Server) Stop() {
	s.grpcServer.Stop()
}

// Search 服务端对文章进行查询。
func (si ServiceImpl) Search(ctx context.Context, searchRequest *SearchRequest) (*SearchResponse, error) {
	response := SearchResponse{}
	articles := []*Article{}
	for _, article := range si.articles {
		if strings.Contains(article.Content, searchRequest.Content) {
			articles = append(articles, article)
		}
	}
	response.Articles = articles
	return &response, nil
}

// Index 服务端对文章进行索引。
func (si ServiceImpl) Index(ctx context.Context, article *Article) (*Empty, error) {
	si.articles[article.Id] = article
	return &Empty{}, nil
}
