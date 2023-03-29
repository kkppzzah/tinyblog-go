// Package interfaces 定义各个接口。
package interfaces

import "encoding/json"

//
const (
	ArticleEventTypeCreate = "create"
	ArticleEventTypeUpdate = "update"
	ArticleEventTypeDelete = "delete"
)

// ArticleInfo 文章信息。
type ArticleInfo struct {
	ID int64 `json:"id"`
}

// ArticleEvent 文章事件。
type ArticleEvent struct {
	EventType   string
	ArticleInfo *ArticleInfo
}

// ArticleEventPublisher 发布用户事件的接口。
type ArticleEventPublisher interface {
	Publish(articleEvent *ArticleEvent) error
}

// ToJSON 输出为JSON。
func (u *ArticleEvent) ToJSON() ([]byte, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// FromJSON 从JSON中解析。
func (u *ArticleEvent) FromJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, u)
}
