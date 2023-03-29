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

// ArticleEventHandler 文章事件处理。
type ArticleEventHandler interface {
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
