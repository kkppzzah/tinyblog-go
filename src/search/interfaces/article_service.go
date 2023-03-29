package interfaces

import (
	"context"

	pb "ppzzl.com/tinyblog-go/search/genproto/article"
)

type ArticleService interface {
	// Publish 发表文章。
	Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error)
	// Update 更新文章。
	Update(ctx context.Context, req *pb.UpdateRequest) (*pb.Empty, error)
	// Delete 删除文章。
	Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error)
	// Get 获取单个文章内容。
	Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error)
	// GetByUser 获取单个用户的多个文章。
	GetByUser(ctx context.Context, req *pb.GetByUserRequest) (*pb.GetByUserResponse, error)
	// GetByIds 获取指定文章id列表的多个文章。
	GetByIds(ctx context.Context, ids []int64) (*pb.GetByIdsResponse, error)
}
