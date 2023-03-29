// main
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "ppzzl.com/tinyblog-go/learn/grpc/hello"
)

// Config 服务端配置。
type Config struct {
	host *string
	port *int
}

func (cfg *Config) initialize() {
	cfg.port = flag.Int("p", 20801, "listening port")
	cfg.host = flag.String("b", "localhost", "binding host")
	flag.Parse()
}

func doSearch(ctx context.Context, c pb.SearchServiceClient, searchContent string) {
	r, err := c.Search(ctx, &pb.SearchRequest{Content: searchContent})
	if err != nil {
		log.Fatalf("failed search article: %v", err)
	}
	log.Printf("search result - %s, %d articles", searchContent, len(r.Articles))
	for _, article := range r.Articles {
		log.Printf("    Id: %d, UserId: %d, Content: %s", article.Id, article.UserId, article.Content)
	}
}

func main() {
	cfg := Config{}
	cfg.initialize()
	address := fmt.Sprintf("%s:%d", *(cfg.host), *(cfg.port))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	log.Printf("connect server %s", address)
	if err != nil {
		log.Fatalf("failed connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = c.Index(ctx, &pb.Article{Id: 1, UserId: 123, Content: "来疑沧海尽成空,万面鼓声中。Hi!"})
	if err != nil {
		log.Fatalf("failed index article: %v", err)
	}
	doSearch(ctx, c, "沧海")
	doSearch(ctx, c, "钱塘自古繁华")
}
